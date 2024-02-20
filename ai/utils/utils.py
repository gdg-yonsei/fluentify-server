import json
import ast


def text2dict(text):
    text = text.strip()
    text = text.replace("\n", "")
    text = text.replace("```json", "")
    text = text.replace("```", "")
    try:
        return ast.literal_eval(text)
    except:
        return json.loads(text) if text else None
