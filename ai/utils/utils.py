import json
import ast


def text2dict(text):
    text = text.strip()
    try:
        return ast.literal_eval(text)
    except:
        return json.loads(text) if text else None
