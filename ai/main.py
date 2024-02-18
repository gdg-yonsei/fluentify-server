from fluentify import Fluentify



if __name__ == "__main__":

    ft = Fluentify()
    
    ## Sample Input ##
    con_input =  {
        "user-answer": "Hmm.. a ship? maybe yellow? I may no",
        "context": "Let's imagine that you are a brave captain of a big ship. You are sailing on the high seas. Suddenly, you see a beautiful sunset. Look at this picture and tell me...",
        "question": "What colors can you see in the sky?",
        "answer": "I see red, orange, yellow, and blue.",
        "img": "1070.jpg"
    }

    pro_input = {
        "user-audio" : "example1.m4a",
        "practice-sentece": "It's autumn now, and the leaves are turning beautiful colors.",
        "tip": "Say 'aw-tum,' not 'ay-tum.'"
    }

    ## Feedback ##
    con_feedback  = ft.ConFeedback(con_input)
    pro_feedback = ft.ProFeedback(pro_input)


    ## Sample Output ##
    print("** Feedback for Contextual Task: ", con_feedback)
    # {'positive-feedback': 'You are very creative! I like your imagination.', 
    #  'negative-feedback': "Let's try to describe what we see in the picture. First, look at the sky. What colors can you see there?",
    #   'enhanced-answer': 'In the sky, I can see yellow, orange, pink, and blue.'}

    print("** Feedback for Pronunciation Task: ", pro_feedback)
    # {'transcription': 'ITS AUTUMN NOW AND THE LEAVES ARE TURNING BEAUTIFUL COLORS', 
    #   'wrong_idx': {'minor': [2, 9], 'major': []}, 
    #   'pronunciation_score': 0.7, 
    #   'decibel': 46.90759735625882, 
    #   'speech_rate': 2.347417840375587, 
    #   'positive-feedback': 'Pronunciation is correct. Keep up the good work!', 
    #   'negative-feedback': ' '}

    
    
    
    
        