package main

import (
	pb "github.com/csh980717/shippy/user-service/proto/auth"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
)

const topic = "auth.created"

type Subscriber struct{}

func main() {
	s := micro.NewService(
		micro.Name("shippy.email"),
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
