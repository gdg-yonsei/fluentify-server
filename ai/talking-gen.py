# Imports the Google Cloud stt library
from google.cloud import texttospeech
from google.api_core import client_options
import stable_whisper
import torch
import librosa
import json
import os 
import cv2
import random
from PIL import Image  
import pandas as pd
from moviepy.editor import VideoFileClip, AudioFileClip, concatenate_videoclips
import time 
from datetime import datetime
# gcloud auth application-default login
# gcloud auth application-default set-quota-project fluentify-412312


class TalkingGenerator:
    def __init__(self):
        self.stt = texttospeech.TextToSpeechClient()
        # self.tts = SpeechClient()
        self.project_id="fluentify-412312"
        self.model = stable_whisper.load_model('base')

        self.data_path = './data'
        self.character_path = f'{self.data_path}/character'
        
        now = datetime.now()
        self.ts  = now.strftime("%Y-%m-%d %H:%M:%S")
        self.dir_path = f'{self.data_path}/{self.ts}'
        
        self.audio_path = self.dir_path + '/output.mp3'
        self.video_path = self.dir_path+ "/output.mp4"
        self.wordoffset_path  = self.dir_path+'/word_offset.json'
        self.output_path = self.dir_path + '/fianal-output.mp4'

        if not os.path.exists(self.dir_path):
            os.makedirs(self.dir_path)

        img_lst = os.listdir(self.character_path)
        self.close_img_lst =   [img for img in img_lst if 'close' in img]
        self.open_img_lst =  [img for img in img_lst if 'open' in img]

        

    def Text2Speech(self,sentence): 
        # Set the text input to be synthesized
        print('Generating Audio file for the sentence: ', sentence)
        synthesis_input = texttospeech.SynthesisInput(text=sentence)

        voice = texttospeech.VoiceSelectionParams(
            language_code="en-US", ssml_gender=texttospeech.SsmlVoiceGender.MALE)

        # Select the type of audio file you want returned
        audio_config = texttospeech.AudioConfig(
            audio_encoding=texttospeech.AudioEncoding.MP3)

        # Perform the text-to-speech request on the text input with the selected
        # voice parameters and audio file type
        response = self.stt.synthesize_speech(
            input=synthesis_input, voice=voice, audio_config=audio_config)

        # The response's audio_content is binary.
        with open(self.audio_path, "wb") as out:
            # Write the response to the output file.
            out.write(response.audio_content)
            print('Audio is generated at: {}'.format(self.audio_path))

        return None
    
    def Speech2Text(self):
        result = self.model.transcribe(self.audio_path)
        
        result.save_as_json(self.wordoffset_path)
        return result
    
    def WordOffset(self):
        with open(self.wordoffset_path, 'r') as f:
            data = json.load(f)
        output = []
        for words in data['segments'] : 
            for word in words['words']:
                output.append({'content': word['word'], 'start': word['start']*100, 'end': word['end']*100})

        for sec in data['nonspeech_sections']:
                output.append({'content': None , 'start': sec['start']*100, 'end': sec['end']*100})
        df = pd.DataFrame(output).sort_values(by=['start'], axis=0)
        df['time'] = df['end'] - df['start']
        self.wordoffset= df
        print('Word Offset is generated at: ', self.wordoffset_path)
        return output
    
    
    def VideoGen(self) : 
        
        img = Image.open(f'{self.character_path}/{self.open_img_lst[0]}')
        width, height = img.size 
        frame_lst = []
        fps = 100 
        close_img = self.close_img_lst[0]
        open_img = self.close_img_lst[0]

        for word in self.wordoffset.iloc:
            ## CLOSE ## 
            if word['content'] == None : 
                # close_img = random.choice(self.close_img_lst)
                for _ in range(int(word['time'])):
                    frame_lst.append(close_img) 

            ## OPEN ## 
            else : 
                # open_img = random.choice(self.open_img_lst)
                for _ in range(int(word['time'])):
                    frame_lst.append(open_img)
                ## switch the image
                if open_img == self.open_img_lst[0]:
                    open_img = self.open_img_lst[1]
                else : 
                    open_img = self.open_img_lst[0]

        codec = cv2.VideoWriter_fourcc(*'mp4v')
        video = cv2.VideoWriter(self.video_path, codec, fps, (width, height)) 
        for frame in frame_lst:  
            video.write(cv2.imread(os.path.join(f'{self.character_path}/{frame}')))
        print('Video is generated at: ', self.video_path)
        cv2.destroyAllWindows()  
        video.release()


    def MergeAudioVideo(self):
        video_clip = VideoFileClip(self.video_path)
        audio_clip = AudioFileClip(self.audio_path)
        final_clip = video_clip.set_audio(audio_clip)
        final_clip.write_videofile(self.output_path)


    def Text2TalkingVideo(self,sentence):
        self.Text2Speech(sentence)
        self.Speech2Text()
        self.WordOffset()
        self.VideoGen()
        self.MergeAudioVideo()
        print('Talking Video is generated at: ', self.output_path)

if __name__ == "__main__":
    tg = TalkingGenerator()
    sentence ="The most important thing to keep in mind when designing this car would be to make sure that it is made of a metal that the superhero can control with his mind."
    #  sentence =  "Let's imagine that you are a brave captain of a big ship. You are sailing on the high seas. Suddenly, you see a beautiful sunset. Look at this picture and tell me..."
    time_  = time.time()
    tg.Text2TalkingVideo(sentence)
    time_  = time.time() - time_
    print(time_)
    
    