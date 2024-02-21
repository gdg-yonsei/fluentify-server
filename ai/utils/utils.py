import ast
import json


def text2dict(text):
    text = text.strip()
    text = text.replace("\n", "")
    text = text.replace("```json", "")
    text = text.replace("```", "")
    try:
        return ast.literal_eval(text)
    except:
        try:
            return json.loads(text)
        except:
            text = text.split('positive-feedback')[-1]
            if 'enhanced-answer' in text:
                pos, neg_w_enh = text.split('negative-feedback')
                neg, enh = neg_w_enh.split('enhanced-answer')
                pos, neg, enh = clean(pos), clean(neg), clean(enh)
                return {"positive-feedback": pos, "negative-feedback": neg, "enhanced-answer": enh}
            else:
                pos, neg = text.split('negative-feedback')
                pos, neg = clean(pos), clean(neg)
                return {"positive-feedback": pos, "negative-feedback": neg}


def clean(text):
    text = text.strip()
    if text.startswith('"') or text.startswith("'"):
        text = text[1:]
    if text.endswith('"') or text.endswith("'"):
        text = text[:-1]

    if ":" in text:
        text = text.replace(":", "")
    if "," in text:
        text = text.replace(",", "")
    if "{" in text:
        text = text.replace("{", "")
    if "}" in text:
        text = text.replace("}", "")

    if "\'s" in text:
        text = text.replace("\'s", " 's")
    if "\'m" in text:
        text = text.replace("\'s", " 'm")

    text = text.strip()

    if text.startswith('"') or text.startswith("'"):
        text = text[1:]
    if text.endswith('"') or text.endswith("'"):
        text = text[:-1]

    text = text.strip()
    return text
