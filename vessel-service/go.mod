module github.com/csh980717/shippy/vessel-service

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
