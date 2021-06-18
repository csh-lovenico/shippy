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

	s := micro.NewService(
		micro.Name("user-service"),
		micro.Version("latest"))
	s.Init()
	pb.RegisterUserServiceHandler(s.Server(), &service{
		repo: repo,
	})
	if err := s.Run(); err != nil {
		log.Println(err)
	}

}
