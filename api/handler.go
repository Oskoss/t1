package api

import (
	"log"

	"github.com/oskoss/t1/datastore"

	"github.com/oskoss/t1/controllers"
	"golang.org/x/net/context"
)

// DeploymentService will handle all Deployment actions
type DeploymentService struct {
}

// Create generates response to a Create request
func (s *DeploymentService) Create(ctx context.Context, in *CreateDeploymentRequest) (*CreateDeploymentResponse, error) {
	log.Printf("Received create request...\n")
	d := controllers.Deployment{Name: in.Name, Created: true}
	r := datastore.Redis{Database: 0, FQDN: "localhost", Password: "", Port: "6379"}
	err := datastore.StoreDeployment(r, d)
	if err != nil {
		log.Printf("Error %+v creating deployment", err)
		return nil, err
	}
	return &CreateDeploymentResponse{}, nil
}

// Stage generates response to a Stage request
func (s *DeploymentService) Stage(ctx context.Context, in *StageDeploymentRequest) (*StageDeploymentResponse, error) {
	log.Printf("Received stage request! Noop for now...will docker pull at some point!\n")
	return &StageDeploymentResponse{}, nil
}

// Status generates response to a Status request
func (s *DeploymentService) Status(ctx context.Context, in *StatusDeploymentRequest) (*StatusDeploymentResponse, error) {
	log.Printf("Received status request...\n")
	d := controllers.Deployment{Name: in.Name}
	r := datastore.Redis{Database: 0, FQDN: "localhost", Password: "", Port: "6379"}
	currentDeployment, err := datastore.StatusDeployment(r, d)
	if err != nil {
		log.Printf("Error %+v getting deployment status", err)
		return nil, err
	}
	log.Printf("Deployment Status : %+v", *currentDeployment)
	return &StatusDeploymentResponse{}, nil
}

// Remove generates response to a Remove request
func (s *DeploymentService) Remove(ctx context.Context, in *RemoveDeploymentRequest) (*RemoveDeploymentResponse, error) {
	log.Printf("Received remove request...\n")
	d := controllers.Deployment{Name: in.Name, Created: true}
	r := datastore.Redis{Database: 0, FQDN: "localhost", Password: "", Port: "6379"}
	err := datastore.RemoveDeployment(r, d)
	if err != nil {
		log.Printf("Error %+v deleting deployment", err)
		return nil, err
	}
	return &RemoveDeploymentResponse{}, nil
}
