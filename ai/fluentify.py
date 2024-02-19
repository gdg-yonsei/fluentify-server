import math
import os

import librosa
import numpy as np
import stable_whisper
import torch
import vertexai
import yaml
from evaluate import load
from transformers import AutoTokenizer, AutoFeatureExtractor, AutoModelForCTC
from vertexai.preview.generative_models import GenerativeModel, Part

from utils.utils import text2dict
from utils.word_process import get_incorrect_idx


# gcloud auth application-default login
# gcloud auth application-default set-quota-project fluentify-412312

class Fluentify:
    def __init__(self):
        vertexai.init(project="fluentify-412312", location="asia-northeast3")
        self.multimodal_model = GenerativeModel("gemini-pro-vision")
        self.lang_model = GenerativeModel("gemini-pro")
        self.current_path = os.path.dirname(__file__)
        self.gcs_path = "gs://fluentify-412312.appspot.com"
        self.shared_path = "./shared-data"
        self.ars_whisper_model = stable_whisper.load_model('base')

        with open(os.path.join(self.current_path, 'data/prompt.yaml'), 'r', encoding='UTF-8') as file:
            self.prompt = yaml.load(file, Loader=yaml.FullLoader)

        self.audio_path = self.shared_path + "/audio"
        self.tokenizer = AutoTokenizer.from_pretrained("facebook/wav2vec2-base-960h")
        self.ars_w2v_model = AutoModelForCTC.from_pretrained("facebook/wav2vec2-base-960h")
        self.feature_extractor = AutoFeatureExtractor.from_pretrained("facebook/wav2vec2-base-960h",
                                                                      sampling_rate=16000)
        self.wer = load("wer")

        self.speech_rate_threshold_h = 2.5
        self.speech_rate_threshold_l = 1.0
        self.decibel_threshold_h = 95
        self.decibel_threshold_l = 45

        self.speech_rate_threshold = (self.speech_rate_threshold_h + self.speech_rate_threshold_l) / 2
        self.decibel_threshold = (self.decibel_threshold_h + self.decibel_threshold_l) / 2

    def GetWer(self, transcription, ground_truth):
        # WER score is from 0 to 1
        wer_score = self.wer.compute(predictions=[transcription], references=[ground_truth])
        incorrect_idx = get_incorrect_idx(ground_truth, transcription)
        incorrect_idx = {"minor": incorrect_idx["substitutions"], "major": incorrect_idx["deletions"]}
        return wer_score, incorrect_idx

    def ASR(self, audio_path):
        ## TODO : Audio File Path ####
        input_audio = librosa.load(audio_path, sr=22000)[0]
        input_values = self.feature_extractor(input_audio, return_tensors="pt", padding="longest").input_values
        rms = librosa.feature.rms(y=input_audio)

        with torch.no_grad():
            logits = self.ars_w2v_model(input_values).logits[0]
            pred_ids = torch.argmax(logits, axis=-1)
        outputs = self.tokenizer.decode(pred_ids, output_word_offsets=True)
        time_offset = self.ars_w2v_model.config.inputs_to_logits_ratio / self.feature_extractor.sampling_rate

        transcription = outputs[0]
        word_offsets = [
            {
                "word": d["word"],
                "start_time": round(d["start_offset"] * time_offset, 2),
                "end_time": round(d["end_offset"] * time_offset, 2),
                "time": round(d["end_offset"] * time_offset - d["start_offset"] * time_offset, 2),
            }
            for d in outputs.word_offsets
        ]

        time = word_offsets[-1]["end_time"]
        speech_rate = len(word_offsets) / word_offsets[-1]["end_time"]
        decibel = 20 * math.log10(np.mean(rms) / 0.00002)
        return transcription, time, speech_rate, decibel

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
        image = Part.from_uri(f"{self.gcs_path}/{input['img']}", mime_type="image/jpeg")
        response = self.ars_whisper_model.transcribe(f"{self.audio_path}/{input['user-audio']}")

        output = {"transcription": response.text}
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
        ground_truth = input["practice-sentence"]
        transcription, time, speech_rate, decibel = self.ASR(f"{self.audio_path}/{input['user-audio']}")
        wer_score, incorrect_idx = self.GetWer(transcription.upper(), ground_truth.upper())

        sentence_lst = ground_truth.split(" ")

        pronunciation_score = self.ScoreMap(1 - wer_score, 1)
        speed_socre = self.ScoreMap(speech_rate, self.speech_rate_threshold)
        volume_score = self.ScoreMap(decibel, self.decibel_threshold)

        output = {
            "transcription": transcription,
            "incorrect_indexes": incorrect_idx['major'],
            "pronunciation_score": pronunciation_score,  # higher the better
            "voulume_score": volume_score,  # higher the better
            "speed_socre": speed_socre,
            "decibel": decibel,
            "speech_rate": speech_rate
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
