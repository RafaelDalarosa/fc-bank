gen:
	protoc --proto_path=infra/grpc infra/grpc/protofile/*.proto --go_out=infra/grpc --go-grpc_out=infra/grpc