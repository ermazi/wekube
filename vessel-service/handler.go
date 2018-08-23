package main

import (
	"gopkg.in/mgo.v2"
	pb "github.com/wekube/vessel-service/proto/vessel"
	"context"
	)

type service struct {
	session *mgo.Session
}


func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) Create(ctx context.Context, req *pb.Vessel,res *pb.Response ) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error   {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)

	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}