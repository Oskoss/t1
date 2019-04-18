package main

import (
	"log"

	"github.com/oskoss/t1/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewDeploymentClient(conn)
	_, err = c.Create(context.Background(), &api.CreateDeploymentRequest{Name: "pbdeployment"})
	if err != nil {
		log.Fatalf("Error when calling Create: %s", err)
	}
	_, err = c.Stage(context.Background(), &api.StageDeploymentRequest{Image: "pbimage"})
	if err != nil {
		log.Fatalf("Error when calling Stage: %s", err)
	}
	_, err = c.Status(context.Background(), &api.StatusDeploymentRequest{Name: "pbdeployment"})
	if err != nil {
		log.Fatalf("Error when calling Status: %s", err)
	}
	_, err = c.Remove(context.Background(), &api.RemoveDeploymentRequest{Name: "pbdeployment"})
	if err != nil {
		log.Fatalf("Error when calling Remove: %s", err)
	}
	_, err = c.Status(context.Background(), &api.StatusDeploymentRequest{Name: "pbdeployment"})
	if err != nil {
		log.Fatalf("Error when calling Status: %s", err)
	}

	log.Printf("Tests Successful")
}
