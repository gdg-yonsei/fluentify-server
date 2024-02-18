## Data Generation ##

**1) 문장 구사** : ```fluentify.ProDataGen(iterarion_num=30)```

   
```data/con-data.json``` : 문장구사평가를 위한 데이터 (```Gemini-pro-vision```)

```data/con-img-pool.json```  : 문장구사평가 데이터 생성에 사용된 [이미지 데이터](https://huggingface.co/datasets/ehristoforu/dalle-3-images) (```Dalle-3``` ) 

Example 
```js

[
    {
        "context": "Let's imagine that you are a brave captain of a big ship. You are sailing on the high seas. Suddenly, you see a beautiful sunset. Look at this picture and tell me...",
        "question": "What colors can you see in the sky?",
        "answer": "I see red, orange, yellow, and blue.",
        "img": "1070.jpg"
    },
   ...
]
```
**2) 발음** : ```fluentify.ConDataGen(iterarion_num=30)```


```data/pro-data.json``` : 발음평가를 위한 데이터 (```Gemini-pro```)

```data/pro-topic-pool.json```  : 발음평가 데이터 생성에 사용된 주제 데이터 (```Gemini-pro```)

Example 
```js
[
    {
        "practice-sentence": "I love to mix baking soda and vinegar together to create a fizzy experiment.",
        "tip": "Remember to say 'mix' with your lips together and 'fizzy' with a big smile."
    },
  ...
]
```



## Feedback Generation ##

**1) 문장 구사** : ```fluentify.ConFeedback(con_input)```


Input
```js
{
    "user-answer": "Hmm.. a ship? maybe yellow? I may no",
    "context": "Let's imagine that you are a brave captain of a big ship. You are sailing on the high seas. Suddenly, you see a beautiful sunset. Look at this picture and tell me...",
    "question": "What colors can you see in the sky?",
    "answer": "I see red, orange, yellow, and blue.",
    "img": "1070.jpg"
}
```
Output
```js
{
    'positive-feedback': 'You are very creative! I like your imagination.', 
    'negative-feedback': "Let's try to describe what we see in the picture. First, look at the sky. What colors can you see there?",
    'enhanced-answer': 'In the sky, I can see yellow, orange, pink, and blue.'
}
```


**2) 발음** : ```fluentify.ProFeedback(pro_input)```

Input 
```js
{
      "user-audio" : "example1.m4a",
      "practice-sentece": "It's autumn now, and the leaves are turning beautiful colors.",
      "tip": "Say 'aw-tum,' not 'ay-tum.'"
}
```
Output
```js
{
    'transcription': 'ITS AUTUMN NOW AND THE LEAVES ARE TURNING BEAUTIFUL COLORS', 
    'wrong_idx': {'minor': [2, 9], 'major': []}, 
    'pronunciation_score': 0.7, 
    'decibel': 46.90759735625882, 
    'speech_rate': 2.347417840375587, 
    'positive-feedback': 'Pronunciation is correct. Keep up the good work!', 
    'negative-feedback': ' '
}
```
----
