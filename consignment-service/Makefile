build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/csh980717/shippy/consignment-service proto/consignment/consignment.proto

buildDocker:
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service