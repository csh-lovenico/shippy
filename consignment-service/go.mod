module github.com/csh980717/shippy/consignment-service

go 1.16

require (
	github.com/csh980717/shippy/user-service v0.0.0-20210712084757-785e8d29ab31 // indirect
	github.com/csh980717/shippy/vessel-service v0.0.0-20210616041555-71b0de97bd56
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	google.golang.org/genproto v0.0.0-20210614182748-5b3b54cad159 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
