build-vessel:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/csh980717/shippy/vessel-service proto/vessel/vessel.proto

build-docker-vessel:
	docker build -t vessel-service .