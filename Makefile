LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	$(LOCAL_BIN)/golangci-lint cache clean
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml



install-deps:
	export GOBIN=/home/dnl/chat/auth/bin && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-api


generate-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path=./api/proto/auth_v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	./api/proto/auth_v1/auth.proto ./api/proto/auth_v1/user.proto


build:
	GOOS=linux GOARCH=amd64 go build -o auth_linux cmd/main.go




