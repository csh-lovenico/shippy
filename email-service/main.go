package main

import (
	"encoding/json"
	pb "github.com/csh980717/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	"log"
)

const topic = "user.created"

func main() {
	s := micro.NewService(
		micro.Name("email-service"),
		micro.Version("latest"))
	s.Init()
	pubSub := s.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatalf("broker connect error: %v\n", err)
	}

	_, err := pubSub.Subscribe(topic, func(pub broker.Event) error {
		var user *pb.User
		if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
			return err
		}
		log.Printf("[Create User]: %v\n", user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Printf("sub error: %v\n", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalf("server run error: %v\n", err)
	}
}

func sendEmail(user *pb.User) error {
	log.Printf("[SENDING A EMAIL TO %s...]", user.Name)
	return nil
}
