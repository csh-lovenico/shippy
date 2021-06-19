package main

import (
	"context"
	"encoding/json"
	"errors"
	pb "github.com/csh980717/shippy/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/v2/config/cmd"
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

	client := pb.NewShippingServiceClient("consignment-service", microclient.DefaultClient)
	file := defaultFilename
	var token string
	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecting file and token. "))
	}
	file = os.Args[1]
	token = os.Args[2]
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
