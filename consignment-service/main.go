package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/wekube/consignment-service/proto/consignment"
	vesselProto "github.com/wekube/vessel-service/proto/vessel"
	"log"
	"os"
)




const (
	defaultHost = "localhost:27017"
)

func main() {
	// get host from env
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// 设置gRPC服务器
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{vesselClient, session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
