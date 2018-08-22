package main

import (
	vesselProto "github.com/wekube/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	pb "github.com/wekube/consignment-service/proto/consignment"
	"log"
	"gopkg.in/mgo.v2"
)

type service struct {
	vesselClient vesselProto.VesselServiceClient
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	})

	if err != nil {
		return err
	}

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	req.VesselId = vesselResponse.Vessel.Id

	// 保存我们的consignment
	err = repo.Create(req)

	if err != nil {
		return err
	}

	res.Created = true
	// 返回的数据也要符合proto中定义的数据结构
	return nil
}

// GetConsignment
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}