package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	"sync"

	pb "github.com/csh980717/shippy/consignment-service/proto/consignment"
)

type repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	mu           sync.Mutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	resp = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments := s.repo.GetAll()
	resp = &pb.Response{Consignments: consignments}
	return nil
}

func main() {
	repo := &Repository{}
	s := micro.NewService(
		micro.Name("consignment"),
		micro.Version("latest"))
	s.Init()
	pb.RegisterShippingServiceHandler(s.Server(), &service{repo})
	if err := s.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
