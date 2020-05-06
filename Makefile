dev:
	go run *.go

build:
	CGO_ENABLED=0 go build -o bin/protobuf

run: build
	./bin/protobuf

prod:
	CGO_ENABLED=0 go build -o bin/protobuf
	upx bin/protobuf

test:
	go test -race -cover ./...

update:
	go get ./...
	go mod tidy
	
pb:
	@protoc --proto_path=pb pb/*.proto --go_out=plugins=grpc:pb

.PHONY: dev build run prod test upadte pb
