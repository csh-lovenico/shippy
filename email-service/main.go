package main

import (
	pb "github.com/csh980717/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/net/context"
	"log"
)

const topic = "user.created"

type Subscriber struct{}

func main() {
	s := micro.NewService(
		micro.Name("email-service"),
		micro.Version("latest"))
	s.Init()
	micro.RegisterSubscriber(topic, s.Server(), new(Subscriber))

	if err := s.Run(); err != nil {
		log.Fatalf("server run error: %v\n", err)
	}
}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("[Picked up a new message]")
	log.Println("[Sending email to]:", user.Name)
	return nil
}
