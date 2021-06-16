package main

import (
	pb "github.com/csh980717/shippy/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type vesselService struct {
	session *mgo.Session
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	defer s.GetRepo().Close()
	vessel, err := s.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func (s *vesselService) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *vesselService) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	defer s.GetRepo().Close()
	if err := s.GetRepo().Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
