module shippy/consignment-cli

go 1.16

require (
	github.com/csh980717/shippy/consignment-service v0.0.0-20210615030706-5eb07286a9d0
	google.golang.org/grpc v1.38.0
)

replace github.com/csh980717/shippy/consignment-service => ../consignment-service
