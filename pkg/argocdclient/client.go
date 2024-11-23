package argocdclient

import (
	"context"
	"errors"
	"log"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

const OPEN_CONNECTION_FAILED_MESSAGE string = "Failed to open a connection to ArgoCD server: %v"

type ArgoCDClient struct {
	client apiclient.Client
	ctx    context.Context
}

func NewArgoCDClient(serverAddr, authToken string, insecure bool, caCertPath string) (*ArgoCDClient, error) {

	clientOpts := &apiclient.ClientOptions{
		ServerAddr: serverAddr,
		AuthToken:  authToken,
		Insecure:   insecure,
		GRPCWeb:    false,
		PlainText:  false,
	}

	client, err := apiclient.NewClient(clientOpts)
	if err != nil {
		log.Fatalf("Failed to create ArgoCD client: %v", err)
		return nil, err
	}

	return &ArgoCDClient{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (c *ArgoCDClient) ListApplications() (*v1alpha1.ApplicationList, error) {

	conn, appClient, err := c.client.NewApplicationClient()
	if err != nil {
		log.Fatalf("Failed to open a connection to ArgoCD server: %v", err)
		return nil, err
	}
	defer conn.Close()

	apps, err := appClient.List(c.ctx, &application.ApplicationQuery{})
	if err != nil {
		log.Printf("Failed to get all applications: %v", err)
		return nil, err
	}

	return apps, nil
}

func (c *ArgoCDClient) CreateApplication(application *application.ApplicationCreateRequest) (*v1alpha1.Application, error) {

	if application == nil {
		return nil, errors.New("application must be defined")
	}

	conn, appClient, err := c.client.NewApplicationClient()
	if err != nil {
		log.Fatalf("Failed to open a connection to ArgoCD server: %v", err)
		return nil, err
	}
	defer conn.Close()

	applicationCreated, err := appClient.Create(c.ctx, application)
	if err != nil {
		log.Fatalf("Application can not be created: %v", err)
		return nil, err
	}
	return applicationCreated, nil
}

func (c *ArgoCDClient) GetApplication(query application.ApplicationQuery) (*v1alpha1.Application, error) {

	if isEmpty(query) {
		return nil, errors.New("Application name parameter must be defined")
	}

	conn, appClient, err := c.client.NewApplicationClient()
	if err != nil {
		log.Fatalf(OPEN_CONNECTION_FAILED_MESSAGE, err)
		return nil, err
	}
	defer conn.Close()

	application, err := appClient.Get(c.ctx, &query)
	if err != nil {
		log.Fatalf("Application not found with query: %v", query)
		return nil, err
	}
	return application, nil
}

func isEmpty(query application.ApplicationQuery) bool {
	return query.Name == nil
}
