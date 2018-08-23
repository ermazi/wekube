package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/wekube/vessel-service/proto/vessel"
	"os"
	"log"
)

const (
	defaultHost = "localhost:27017"
)

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
		host = ""
	}
	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Cannot connect to mongodb with url:%v, err:%v", host, err)
	}

	createDummyData(&VesselRepository{session.Copy()})

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
