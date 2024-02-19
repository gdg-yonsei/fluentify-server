proto:
	python -m grpc_tools.protoc -I../idl/proto --python_out=. --pyi_out=. --grpc_python_out=. ../idl/proto/internal.proto

clean:
	rm internal_pb2.py internal_pb2_grpc.py internal_pb2.pyi