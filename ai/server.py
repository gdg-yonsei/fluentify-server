from concurrent import futures
import logging

import grpc
import internal_pb2
import internal_pb2_grpc


class HelloService(internal_pb2_grpc.HelloServiceServicer):
    def Hello(self, request, context):
        print(request.name)
        return internal_pb2.HelloResponse(message='Hello %s' % request.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    internal_pb2_grpc.add_HelloServiceServicer_to_server(HelloService(), server)

    # 8081번 포트로 서버 열기
    server.add_insecure_port('[::]:8081')
    print("server on port 8081")
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
