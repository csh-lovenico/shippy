package main

import (
	pb "github.com/csh980717/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/csh980717/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const defaultHost = "ksks.bokurano.live:27017"

//type Repository struct {
//	consignments []*pb.Consignment
//}
//
//func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
//	updated := append(repo.consignments, consignment)
//	repo.consignments = updated
//	return consignment, nil
//}
//
//func (repo *Repository) GetAll() []*pb.Consignment {
//	return repo.consignments
//}
//
//type service struct {
//	repo         Repository
//	vesselClient vesselProto.VesselServiceClient
//}
//
//func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
//	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{MaxWeight: req.Weight, Capacity: int32(len(req.Containers))})
//	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
//	if err != nil {
//		return err
//	}
//	req.VesselId = vesselResponse.Vessel.Id
//	consignment, err := s.repo.Create(req)
//	if err != nil {
//		return err
//	}
//	resp.Created = true
//	resp.Consignment = consignment
//	return nil
//}
//
//func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
//	consignments := s.repo.GetAll()
//	resp.Consignments = consignments
//	return nil
//}

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
