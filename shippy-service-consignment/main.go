package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"

	pb "github.com/galahade/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/galahade/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)
	// Init will parse the command line flags.
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	log.Println("Start to connect to shippy.service.vessel service.")
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)

	}
	log.Println("Start to listen as service")
	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
