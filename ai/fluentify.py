import os
import requests
import yaml 
import json
import vertexai
import numpy as np 
from vertexai.preview.generative_models import GenerativeModel, Part
import random
import ast

import torch
from transformers import AutoTokenizer, AutoFeatureExtractor, AutoModelForCTC
import torch
from jiwer import wer
import os 
import json
import math
import torch
import librosa
from evaluate import load
from jiwer import compute_measures
import numpy as np
from utils.word_process import get_wrong_idx


# gcloud auth application-default login
# gcloud auth application-default set-quota-project fluentify-412312

class Fluentify:
    def __init__(self):
        vertexai.init(project="fluentify-412312", location="asia-northeast3")
        self.multimodal_model = GenerativeModel("gemini-pro-vision")
        self.lang_model = GenerativeModel("gemini-pro")
        self.current_path = os.path.dirname(__file__)
        self.gcs_path = "gs://fluentify-412312.appspot.com"
        with open(os.path.join(self.current_path,'data/prompt.yaml'), 'r', encoding='UTF-8') as file:
            self.prompt = yaml.load(file, Loader=yaml.FullLoader)

        self.audio_path = "./data/audio"
        self.tokenizer =  AutoTokenizer.from_pretrained("facebook/wav2vec2-base-960h")
        self.model = AutoModelForCTC.from_pretrained("facebook/wav2vec2-base-960h")
        self.feature_extractor = AutoFeatureExtractor.from_pretrained("facebook/wav2vec2-base-960h")
        self.wer = load("wer")

        self.speech_rate_threshol_h = 2.5
        self.speech_rate_threshol_l = 1.0
        self.decibel_threshold_h = 95
        self.decibel_threshold_l = 45
        
    def GetWer(self, transcription, ground_truth):
        # WER score is from 0 to 1
        wer_score = self.wer.compute(predictions=[transcription], references=[ground_truth])
        wrong_idx = get_wrong_idx(ground_truth, transcription)
        wrong_idx = {"minor":wrong_idx["substitutions"], "major":wrong_idx["deletions"]}
        return wer_score, wrong_idx

    def ASR(self, audio):
        ## TODO : Audio File Path ####
        input_audio, _ = librosa.load(f"./data/audio/{audio}")

        input_values = self.feature_extractor(input_audio, return_tensors="pt", padding="longest").input_values
        rms = librosa.feature.rms(y=input_audio)

        with torch.no_grad():
            logits = self.model(input_values).logits[0]
            pred_ids = torch.argmax(logits, axis=-1)
        outputs = self.tokenizer.decode(pred_ids, output_word_offsets=True)
        time_offset = self.model.config.inputs_to_logits_ratio / self.feature_extractor.sampling_rate

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

    def ConFeedback(self, input):
        image = Part.from_uri(f"{self.gcs_path}/img/{input['img']}", mime_type="image/jpeg")

        #### TODO : Feedback focusing more on sentence formulation ##### 
        prompt = f"{self.prompt['con-feedback']}".format(context =input["context"] ,question=input["question"], answer=input["answer"], user_answer=input["user-answer"])
        response = self.multimodal_model.generate_content([prompt, image])
        # print(response.text)
        try : 
            return ast.literal_eval((response.text).strip())
        except:
            return None 
    
    def ProFeedback(self, input):
        ground_truth = input["practice-sentece"]
        transcription, time, speech_rate, decibel = self.ASR(input["user-audio"])

        wer_score, wrong_idx = self.GetWer(transcription.upper(), ground_truth.upper())

        pronunciation_score = 1-wer_score
        sentence_lst = ground_truth.split(" ")
        
        output = {
            "transcription": transcription,
            "wrong_idx": wrong_idx,
            "pronunciation_score": pronunciation_score, # higher the better
            "decibel": decibel,
            "speech_rate": speech_rate,
        }

        feedback ={
            "speech_rate":"You are speaking in a good speed.",
            "decibel":"You are speaking in a good volume.",
            "wrong_pronunciation": "Nothing"
        }

        if self.speech_rate_threshol_h < output['speech_rate'] :
            feedback['speech_rate'] = "You are speaking too fast."
        if self.speech_rate_threshol_l > output['speech_rate'] :
            feedback['speech_rate'] = "You are speaking too slow."

        if self.decibel_threshold_h < output['decibel'] :
            feedback['decibel'] = "You are speaking too loud." 
        if self.decibel_threshold_l > output['decibel'] :
            feedback['decibel'] = "You are speaking too quiet."

        ## TODO : Decide whether to use only major errors or only minor errors ## 
        # only if there is major wrong pronunciation 
        if len(output["wrong_idx"]['major']) > 0:
            for idx in output['major']:
                feedback['wrong_pronunciation'] = f"Your pronunciation for {sentence_lst[idx]} is not correct"


        prompt = f"{self.prompt['pro-feedback']}".format(user_answer = output['transcription'] , 
                                                         speech_rate =feedback['speech_rate'], 
                                                         decibel_level = feedback['decibel'], 
                                                         wrong_pronunciation = feedback['wrong_pronunciation'])
        response =self.lang_model.generate_content(prompt)
        # print(response.text)
        try : 
            feedbaack =  ast.literal_eval((response.text).strip())
            output.update(feedbaack)
            return output
        except:
            return None 
    

