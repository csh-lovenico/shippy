package main

import (
	"context"
	"errors"
	pb "github.com/csh980717/shippy/consignment-service/proto/consignment"
	"github.com/csh980717/shippy/user-service/proto/auth"
	vesselProto "github.com/csh980717/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"os"
)

const defaultHost = "localhost:27017"

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	s := micro.NewService(
		micro.Name("shippy.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper))
	vesselClient := vesselProto.NewVesselServiceClient("shippy.vessel", s.Client())
	s.Init()
	pb.RegisterShippingServiceHandler(s.Server(), &service{session, vesselClient})
	if err := s.Run(); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, rsp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["token"]
		log.Println("Authenticating with token: ", token)
		authClient := auth.NewAuthClient("shippy.auth", microclient.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &auth.Token{Token: token})
		if err != nil {
			return err
		}
		err = fn(ctx, req, rsp)
		return err
	}
}
