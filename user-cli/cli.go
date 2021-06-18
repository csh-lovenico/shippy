package main

import (
	pb "github.com/csh980717/shippy/user-service/proto/user"
	"github.com/micro/cli"
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
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company"},
		),
	)
	service.Init(
		micro.Action(func(ctx *cli.Context) {
			name := ctx.String("name")
			email := ctx.String("email")
			password := ctx.String("password")
			company := ctx.String("company")
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %v", r.User.Id)
			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}
			os.Exit(0)
		}))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
