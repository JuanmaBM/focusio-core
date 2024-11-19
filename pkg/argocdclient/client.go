package argocdclient

import (
	"context"
	"errors"
	"log"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

type ArgoCDClient struct {
	client apiclient.Client
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

	return &ArgoCDClient{client: client}, nil
}

func (c *ArgoCDClient) ListApplications() (*v1alpha1.ApplicationList, error) {

	conn, appClient, err := c.client.NewApplicationClient()
	if err != nil {
		log.Fatalf("Failed to open a connection to ArgoCD server: %v", err)
		return nil, err
	}
	defer conn.Close()

	ctx := context.Background()
	apps, err := appClient.List(ctx, &application.ApplicationQuery{})
	if err != nil {
		log.Printf("Failed to get all applications: %v", err)
		return nil, err
	}

	return apps, nil
}

func (c *ArgoCDClient) createApplication(application *application.ApplicationCreateRequest) (*v1alpha1.Application, error) {

	if application == nil {
		return nil, errors.New("application must be defined")
	}

	conn, appClient, err := c.client.NewApplicationClient()
	if err != nil {
		log.Fatalf("Failed to open a connection to ArgoCD server: %v", err)
		return nil, err
	}
	defer conn.Close()

	ctx := context.Background()
	applicationCreated, err := appClient.Create(ctx, application)
	if err != nil {
		log.Fatalf("Application can not be created: %v", err)
		return nil, err
	}
	return applicationCreated, nil
}
