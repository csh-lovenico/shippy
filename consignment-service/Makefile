build-consignment:
	protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto

build-docker-consignment:
	docker build -t consignment-service .

run-consignment:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service