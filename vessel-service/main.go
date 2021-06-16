package main

import (
	pb "github.com/csh980717/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const defaultHost = "localhost:27017"

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

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

	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)

	service := micro.NewService(
		micro.Name("vessel-service"))
	service.Init()
	pb.RegisterVesselServiceHandler(service.Server(), &vesselService{session})
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
