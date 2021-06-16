build-vessel:
	protoc --proto_path=. --go_out=. --micro_out=. proto/vessel/vessel.proto

build-docker-vessel:
	docker build -t vessel-service .