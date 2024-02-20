from concurrent import futures
import logging

import grpc
import internal_pb2
import internal_pb2_grpc
import fluentify


class PronunciationFeedbackService(internal_pb2_grpc.PronunciationFeedbackServiceServicer):
    def __init__(self):
        self.ft = fluentify.Fluentify()

    def PronunciationFeedback(self, request, context):
        pro_input = {
            "user-audio": request.audio_path,
            "practice-sentence": request.sentence,
            "tip": request.tip
        }
        feedback = self.ft.ProFeedback(pro_input)
        print(feedback)

        return internal_pb2.PronunciationFeedbackResponse(
            transcript=feedback["transcription"],
            incorrect_indexes=feedback["incorrect_indexes"],
            pronunciation_score=feedback["pronunciation_score"],
            decibel=feedback["decibel"],
            speech_rate=feedback["speech_rate"],
            positive_feedback=feedback["positive_feedback"],
            negative_feedback=feedback["negative_feedback"]
        )


class CommunicationFeedbackService(internal_pb2_grpc.CommunicationFeedbackServiceServicer):
    def __init__(self):
        self.ft = fluentify.Fluentify()

    def CommunicationFeedback(self, request, context):
        com_input = {
            "user-audio": request.audio_path,
            "context": request.context,
            "question": request.question,
            "answer": request.expected_answer,
            "img": request.img_path,
        }
        feedback = self.ft.ComFeedback(com_input)
        print(feedback)

        return internal_pb2.CommunicationFeedbackResponse(
            positive_feedback=feedback["positive_feedback"],
            negative_feedback=feedback["negative_feedback"],
            enhanced_answer=feedback["enhanced_answer"]
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    internal_pb2_grpc.add_PronunciationFeedbackServiceServicer_to_server(PronunciationFeedbackService(), server)
    internal_pb2_grpc.add_CommunicationFeedbackServiceServicer_to_server(CommunicationFeedbackService(), server)

    # 8081번 포트로 서버 열기
    server.add_insecure_port('[::]:8081')
    print("server on port 8081")
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
