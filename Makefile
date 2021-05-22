generate:
	buf generate --path ./src/protos/coordservice/coordservice.proto

install:
	go get \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/bufbuild/buf/cmd/buf
