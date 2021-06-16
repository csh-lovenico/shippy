module shippy/consignment-cli

go 1.16

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210512092938-c05353c2d58c // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/csh980717/shippy/consignment-service v0.0.0-20210616035644-80f3bc512b63
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/kevinburke/ssh_config v1.1.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	//github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.42 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/nats-io/jwt v1.2.2 // indirect
	github.com/nats-io/nats.go v1.11.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	//google.golang.org/grpc/examples v0.0.0-20210614190250-22c535818725 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/csh980717/shippy/consignment-service => ../consignment-service

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
