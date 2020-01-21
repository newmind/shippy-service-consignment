// shippy-service-consignment/main.go
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/newmind/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/newmind/shippy-service-vessel/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "mongodb://localhost:32773"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
