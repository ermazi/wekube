package main

import (
	pb "github.com/wekube/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
	"log"
)
const (
	dbName = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(specification *pb.Specification) (*pb.Vessel,error)
	Create (*pb.Vessel) error
	Close()
}


type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection  {
	return repo.session.DB(dbName).C(vesselCollection)
}

func (repo *VesselRepository) Close()  {
	repo.session.Close()
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}


func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error)  {
	var vessel *pb.Vessel

	log.Printf("input spec: %v", spec)
	err := repo.collection().Find(bson.M{
		"capacity": bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	log.Printf("result spec: %v", vessel)

	return vessel, nil
}