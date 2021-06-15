module shippy/consignment-cli

go 1.16

require (
	github.com/csh980717/shippy/consignment-service v0.0.0-20210615030706-5eb07286a9d0
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	google.golang.org/grpc v1.38.0
)

replace github.com/csh980717/shippy/consignment-service => ../consignment-service
