build-vessel:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/csh980717/shippy/vessel-service proto/vessel/vessel.proto

build-docker-vessel:
	docker build -t vessel-service .

run-vessel:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service