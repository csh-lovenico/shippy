package main

import (
	pb "github.com/csh980717/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/csh980717/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const defaultHost = "ksks.bokurano.live:27017"

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
		micro.Name("consignment-service"),
		micro.Version("latest"))
	vesselClient := vesselProto.NewVesselServiceClient("vessel-service", s.Client())
	s.Init()
	pb.RegisterShippingServiceHandler(s.Server(), &service{session, vesselClient})
	if err := s.Run(); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
