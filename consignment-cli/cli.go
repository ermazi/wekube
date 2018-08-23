package main

import (
	"encoding/json"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/wekube/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}


func main() {
	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)
	//wg := sync.WaitGroup{}
	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	//for i := 0; i <= 10000 ; i++ {
	//	wg.Add(1)
	//go func(i int, wg *sync.WaitGroup) {

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Printf("Could not create consignment: %v", err)
		panic(err)
	}
	log.Printf("Created: %v", r.Created)
	//		wg.Done()
	//	}(i, &wg)
	//
	//wg.Wait()
	//getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	//if err != nil {
	//	log.Fatalf("Could not list consignments: %v", err)
	//}
	//for _, v := range getAll.Consignments {
	//	log.Println(v)
	//}
}
