build-consignment:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/csh980717/shippy/consignment-service proto/consignment/consignment.proto

build-docker-consignment:
	docker build -t consignment-service .

run-consignment:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service