package main

import (
	pb "github.com/wekube/consignment-service/proto/consignment"

	"gopkg.in/mgo.v2"
)

const (
	dbName = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll()([]*pb.Consignment, error)
	collection() *mgo.Collection
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) collection() *mgo.Collection  {
	return repo.session.DB(dbName).C(consignmentCollection)
}

func (repo *ConsignmentRepository) Close()  {
	repo.session.Close()
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error  {
	return repo.collection().Insert(consignment)
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error)  {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}