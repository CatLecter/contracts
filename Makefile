generate:
	protoc proto/v1/* --go_out=./gen/go/ --go-grpc_out=./gen/go/
