import grpc

from python_part.gprc_start.proto import helloworld_pb2, helloworld_pb2_grpc

if __name__ == '__main__':
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
        print("Greeter client received: " + response.message)

