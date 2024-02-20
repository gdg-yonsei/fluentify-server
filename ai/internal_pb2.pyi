from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class PronunciationFeedbackRequest(_message.Message):
    __slots__ = ("sentence", "audio_path", "tip")
    SENTENCE_FIELD_NUMBER: _ClassVar[int]
    AUDIO_PATH_FIELD_NUMBER: _ClassVar[int]
    TIP_FIELD_NUMBER: _ClassVar[int]
    sentence: str
    audio_path: str
    tip: str
    def __init__(self, sentence: _Optional[str] = ..., audio_path: _Optional[str] = ..., tip: _Optional[str] = ...) -> None: ...

class PronunciationFeedbackResponse(_message.Message):
    __slots__ = ("transcript", "incorrect_indexes", "pronunciation_score", "decibel", "speech_rate", "positive_feedback", "negative_feedback")
    TRANSCRIPT_FIELD_NUMBER: _ClassVar[int]
    INCORRECT_INDEXES_FIELD_NUMBER: _ClassVar[int]
    PRONUNCIATION_SCORE_FIELD_NUMBER: _ClassVar[int]
    DECIBEL_FIELD_NUMBER: _ClassVar[int]
    SPEECH_RATE_FIELD_NUMBER: _ClassVar[int]
    POSITIVE_FEEDBACK_FIELD_NUMBER: _ClassVar[int]
    NEGATIVE_FEEDBACK_FIELD_NUMBER: _ClassVar[int]
    transcript: str
    incorrect_indexes: _containers.RepeatedScalarFieldContainer[int]
    pronunciation_score: float
    decibel: float
    speech_rate: float
    positive_feedback: str
    negative_feedback: str
    def __init__(self, transcript: _Optional[str] = ..., incorrect_indexes: _Optional[_Iterable[int]] = ..., pronunciation_score: _Optional[float] = ..., decibel: _Optional[float] = ..., speech_rate: _Optional[float] = ..., positive_feedback: _Optional[str] = ..., negative_feedback: _Optional[str] = ...) -> None: ...

class CommunicationFeedbackRequest(_message.Message):
    __slots__ = ("context", "question", "expected_answer", "audio_path", "img_path")
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    QUESTION_FIELD_NUMBER: _ClassVar[int]
    EXPECTED_ANSWER_FIELD_NUMBER: _ClassVar[int]
    AUDIO_PATH_FIELD_NUMBER: _ClassVar[int]
    IMG_PATH_FIELD_NUMBER: _ClassVar[int]
    context: str
    question: str
    expected_answer: str
    audio_path: str
    img_path: str
    def __init__(self, context: _Optional[str] = ..., question: _Optional[str] = ..., expected_answer: _Optional[str] = ..., audio_path: _Optional[str] = ..., img_path: _Optional[str] = ...) -> None: ...

class CommunicationFeedbackResponse(_message.Message):
    __slots__ = ("positive_feedback", "negative_feedback", "enhanced_answer")
    POSITIVE_FEEDBACK_FIELD_NUMBER: _ClassVar[int]
    NEGATIVE_FEEDBACK_FIELD_NUMBER: _ClassVar[int]
    ENHANCED_ANSWER_FIELD_NUMBER: _ClassVar[int]
    positive_feedback: str
    negative_feedback: str
    enhanced_answer: str
    def __init__(self, positive_feedback: _Optional[str] = ..., negative_feedback: _Optional[str] = ..., enhanced_answer: _Optional[str] = ...) -> None: ...
