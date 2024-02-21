import math
import os

import librosa
import numpy as np
import stable_whisper
import vertexai
import yaml
from evaluate import load
from vertexai.preview.generative_models import GenerativeModel, Part

from utils.utils import text2dict
from utils.word_process import get_incorrect_idxes


# gcloud auth application-default login
# gcloud auth application-default set-quota-project fluentify-412312

class Fluentify:
    def __init__(self):
        vertexai.init(project="fluentify-412312", location="asia-northeast3")
        self.multimodal_model = GenerativeModel("gemini-pro-vision")
        self.lang_model = GenerativeModel("gemini-pro")
        self.current_path = os.path.dirname(__file__)
        self.gcs_path = "gs://fluentify-412312.appspot.com"
        self.ars_model = stable_whisper.load_model('base')

        with open(os.path.join(self.current_path, 'data/prompt.yaml'), 'r', encoding='UTF-8') as file:
            self.prompt = yaml.load(file, Loader=yaml.FullLoader)

        self.audio_path = "./shared-data/audio"
        # self.tokenizer =  AutoTokenizer.from_pretrained("facebook/wav2vec2-base-960h")
        # self.ars_w2v_model = AutoModelForCTC.from_pretrained("facebook/wav2vec2-base-960h")
        # self.feature_extractor = AutoFeatureExtractor.from_pretrained("facebook/wav2vec2-base-960h", sampling_rate=16000)
        self.wer = load("wer")

        self.speech_rate_threshold_h = 4.5
        self.speech_rate_threshold_l = 2.5
        self.decibel_threshold_h = 95
        self.decibel_threshold_l = 45

        self.speech_rate_threshold = (self.speech_rate_threshold_h + self.speech_rate_threshold_l) / 2
        self.decibel_threshold = (self.decibel_threshold_h + self.decibel_threshold_l) / 2

    def GetWer(self, transcription, ground_truth):
        # WER score is from 0 to 1
        wer_score = self.wer.compute(predictions=[transcription], references=[ground_truth])
        incorrect_idxes = get_incorrect_idxes(ground_truth, transcription)
        self.wer_score = wer_score
        self.incorrect_idxes = incorrect_idxes["substitutions"] + incorrect_idxes["deletions"]

    def GetDecibel(self, audio_path):
        input_audio, _ = librosa.load(audio_path, sr=22000)
        rms = librosa.feature.rms(y=input_audio)
        decibel = 20 * math.log10(np.mean(rms) / 0.00002)
        return decibel

    def EvaluatePro(self, audio_path):
        transcription = self.ars_model.transcribe(audio_path)
        self.pro_transcription = transcription.text

        word_offset = transcription.segments
        start, end = word_offset[0].start, word_offset[-1].end
        duration = end - start

        self.speech_rate = len(self.pro_transcription.split(' ')) / duration
        self.decibel = self.GetDecibel(audio_path)

    def Score2Des(self, output):
        feedback = {
            "speech_rate": "You are speaking in a good speed.",
            "decibel": "You are speaking in a good volume.",
            "incorrect_pronunciation": "Nothing"
        }

        if self.speech_rate_threshold_h < output['speech_rate']:
            feedback['speech_rate'] = "You are speaking too fast."
        if self.speech_rate_threshold_l > output['speech_rate']:
            feedback['speech_rate'] = "You are speaking too slow."

        if self.decibel_threshold_h < output['decibel']:
            feedback['decibel'] = "You are speaking too loud."
        if self.decibel_threshold_l > output['decibel']:
            feedback['decibel'] = "You are speaking too quiet."
        return feedback

    def ScoreMap(self, input, threshold):
        ratio = abs((threshold - input) / threshold)
        if ratio > 0.8:
            return 1
        elif ratio > 0.6:
            return 2
        elif ratio > 0.4:
            return 3
        elif ratio > 0.2:
            return 4
        else:
            return 5

    def ComFeedback(self, input):
        image = Part.from_uri(f"{self.gcs_path}/img/{input['img']}", mime_type="image/jpeg")
        response = self.ars_model.transcribe(f"{self.audio_path}/{input['user-audio']}")
        self.com_trnascription = response.text
        output = {"transcription": self.com_trnascription}

        #### TODO : Feedback focusing more on sentence formulation #####
        prompt = f"{self.prompt['con-feedback']}".format(context=input["context"], question=input["question"],
                                                         answer=input["answer"], user_answer=output['transcription'])
        response = self.multimodal_model.generate_content([prompt, image])
        try:
            dict_output = text2dict(response.text)
            output.update(dict_output)
            return output
        except:
            print(text2dict(response.text))
            print("Error in ComFeedback response")
            return None

    def ProFeedback(self, input):
        ground_truth = input["practice-sentece"]
        self.EvaluatePro(f"{self.audio_path}/{input['user-audio']}")
        self.GetWer(self.pro_transcription.upper(), ground_truth.upper())
        sentence_lst = ground_truth.split(" ")

        pronunciation_score = self.ScoreMap(1 - self.wer_score, 1)
        speed_score = self.ScoreMap(self.speech_rate, self.speech_rate_threshold)
        volume_score = self.ScoreMap(self.decibel, self.decibel_threshold)

        output = {
            "transcription": self.pro_transcription,
            "incorrect_indexes": self.incorrect_idxes,
            "decibel": self.decibel,
            "speech_rate": self.speech_rate,
            "pronunciation_score": pronunciation_score,  # higher the better
            "volume_score": volume_score,  # higher the better
            "speed_score": speed_score
        }

        feedback = self.Score2Des(output)

        ## TODO : Decide whether to use only major errors or only minor errors ##
        # only if there is major incorrect pronunciation
        if len(output["incorrect_indexes"]) > 0:
            for idx in output["incorrect_indexes"]:
                feedback['incorrect_pronunciation'] = f"Your pronunciation for {sentence_lst[idx]} is not correct"

        prompt = f"{self.prompt['pro-feedback']}".format(user_answer=output['transcription'],
                                                         speech_rate=feedback['speech_rate'],
                                                         decibel_level=feedback['decibel'],
                                                         incorrect_pronunciation=feedback['incorrect_pronunciation'])
        response = self.lang_model.generate_content(prompt)

        try:
            dict_output = text2dict(response.text)
            output.update(dict_output)
            return output
        except:
            print(text2dict(response.text))
            print("Error in ProFeedback response")
            return None
