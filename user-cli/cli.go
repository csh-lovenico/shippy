package main

import (
	pb "github.com/csh980717/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {
	cmd.Init()
	client := pb.NewUserServiceClient("user-service", microclient.DefaultClient)
	service := micro.NewService(
		micro.Name("user-service-cli"),
		micro.Version("latest"),
	)
	service.Init()
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"
	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)
	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is: %s \n", authResponse.Token)
	os.Exit(0)
}
