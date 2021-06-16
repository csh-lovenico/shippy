module shippy/consignment-cli

go 1.16

require (
	github.com/csh980717/shippy/consignment-service v0.0.0-20210615030706-5eb07286a9d0
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc/examples v0.0.0-20210614190250-22c535818725 // indirect
)

replace github.com/csh980717/shippy/consignment-service => ../consignment-service

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
