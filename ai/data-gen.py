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

class DataGeneration:
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

    def GenSent(self, topic):
        topic = " ".join(topic) if type(topic) == list else topic
        # print('used topic:', topic)
        # print(self.prompt['gen-sent'])
        prompt = f"{self.prompt['gen-sent']}".format(topic=topic)
        response =self.lang_model.generate_content(prompt)
        response=response.text.replace("```json","")
        response=response.replace("```","")

        try : 
            return ast.literal_eval(response)
        except:
            return None 

    def ImgFilter(self, img):
        image = Part.from_uri(f"{self.gcs_path}/img/{img}", mime_type="image/jpeg")
        prompt = self.prompt['img-filter']
        response = self.multimodal_model.generate_content([prompt, image])
        output = False if  "No" in response.text  else True
        # print(output)
        return output
    
    def GenQA(self, img):
        image = Part.from_uri(f"{self.gcs_path}/img/{img}", mime_type="image/jpeg")
        prompt = self.prompt['gen-qa']
        response = self.multimodal_model.generate_content([prompt, image])
        # print(image, response.text)
        try : 
            return ast.literal_eval((response.text).strip())
        except:
            return None 

    def ConDataGen(self, iterate_num):
        data_path = f"{self.current_path}/data/con-data.json"
        img_pool_path = f"{self.current_path}/data/con-img-pool.json"
        with open(img_pool_path) as f:
            img_pool = json.load(f) 
        with open(data_path) as f:
            con_data = json.load(f) 

        # Generate Questions and Answers for ConText Evaluation
        for i in range(iterate_num):
            img = random.sample(img_pool, 1)
            img = img[0] if type(img) == list else img

            img_pool.remove(img) 
            with open(img_pool_path,'w') as f:
                    json.dump(img_pool, f,indent=4)
            

            filter = self.ImgFilter(img)
            if filter:
                qa = self.GenQA(img) 
                # if parsing is successful
                if qa:
                    qa.update({"img": img}) 
                    print(qa)
                    con_data += [qa]
                    with open(data_path, "w") as f:
                        json.dump(con_data, f,indent=4)
            else:
                print(f"Image {img} is not suitable for the context evaluation.")
        return con_data
        
    def ProDataGen(self, iterate_num):
        data_path = f"{self.current_path}/data/pro-data.json"
        topic_path = f"{self.current_path}/data/pro-topic-pool.json"
        with open(topic_path) as f:
            topic_pool = json.load(f) 
        with open(data_path) as f:
            pro_data = json.load(f) 

        # Generate Sentences and Tips for Pronunciation Evaluation
        for i in range(iterate_num):
            topic = np.random.choice(topic_pool, size=5, replace=False)
            sent = self.GenSent(topic) 
            # if parsing is successful
            if sent  : 
                print(sent)
                pro_data += sent
                with open(data_path, "w") as f:
                    json.dump(pro_data, f,indent=4)
        return pro_data
  