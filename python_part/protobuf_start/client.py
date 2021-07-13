from python_part.protobuf_start.proto import hello_pb2

request = hello_pb2.HelloRequest()
request.name = "bobby"
res_str = request.SerializeToString()
print(res_str)

request2 = hello_pb2.HelloRequest()
request2.ParseFromString(res_str)
print(request2.name)