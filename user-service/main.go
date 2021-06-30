package main

import (
	pb "github.com/csh980717/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db: db}
	tokenService := &TokenService{repo: repo}
	s := micro.NewService(
		micro.Name("user-service"),
		micro.Version("latest"))
	s.Init()
	pubSub := s.Server().Options().Broker
	pb.RegisterUserServiceHandler(s.Server(), &service{
		repo:         repo,
		tokenService: tokenService,
		pubSub:       pubSub,
	})
	if err := s.Run(); err != nil {
		log.Println(err)
	}

}
