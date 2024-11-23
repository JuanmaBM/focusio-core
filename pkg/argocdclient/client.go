package argocdclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

type ArgoCDClient struct {
	client     apiclient.Client
	ctx        context.Context
	tokenLock  sync.Mutex
	username   string
	password   string
	serverAddr string
	authToken  string
	insecure   bool
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func getAuthToken(serverAddr string, username string, password string) (string, error) {
	loginURL := serverAddr + "/api/v1/session"
	loginRequest := LoginRequest{
		Username: username,
		Password: password,
	}

	reqBody, err := json.Marshal(loginRequest)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to login: %s", string(bodyBytes))
	}

	var loginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return "", err
	}

	return loginResponse.Token, nil
}

func createArgcdClient(serverAddr string, authToken string, insecure bool) (apiclient.Client, error) {

	clientOpts := &apiclient.ClientOptions{
		ServerAddr: serverAddr,
		AuthToken:  authToken,
		Insecure:   insecure,
		GRPCWeb:    false,
		PlainText:  false,
	}

	client, err := apiclient.NewClient(clientOpts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewArgoCDClient(serverAddr string, username string, password string, insecure bool, caCertPath string) (*ArgoCDClient, error) {

	authToken, err := getAuthToken(serverAddr, username, password)
	if err != nil {
		log.Fatalf("Client can't get Authorization Token from ArgoCD with the crendetials provided")
		return nil, err
	}

	client, err := createArgcdClient(serverAddr, authToken, insecure)
	if err != nil {
		log.Fatalf("Failed to create ArgoCD client: %v", err)
		return nil, err
	}

	return &ArgoCDClient{
		client:     client,
		ctx:        context.Background(),
		username:   username,
		password:   password,
		serverAddr: serverAddr,
		authToken:  authToken,
		insecure:   insecure,
	}, nil
}

func (c *ArgoCDClient) DoRequestWithRetry(requestFunc func(appClient application.ApplicationServiceClient) error) error {

	conn, appClient, err := c.client.NewApplicationClient()
	if err != nil {
		return fmt.Errorf("failed to open a connection to ArgoCD server: %v", err)
	}
	defer conn.Close()

	err = requestFunc(appClient)

	if err != nil && isUnauthorized(err) {

		c.tokenLock.Lock()
		defer c.tokenLock.Unlock()

		authToken, err := getAuthToken(c.serverAddr, c.username, c.password)
		if err != nil {
			return fmt.Errorf("error renewing auth token: %v", err)
		}

		c.authToken = authToken
		c.client, err = createArgcdClient(c.serverAddr, authToken, c.insecure)
		if err != nil {
			return fmt.Errorf("error recreating ArgoCD client with new token: %v", err)
		}

		err = requestFunc(appClient)
	}

	return err
}

func (c *ArgoCDClient) ListApplications() (*v1alpha1.ApplicationList, error) {

	var apps *v1alpha1.ApplicationList
	err := c.DoRequestWithRetry(func(appClient application.ApplicationServiceClient) error {
		appList, err := appClient.List(c.ctx, &application.ApplicationQuery{})
		if err != nil {
			return fmt.Errorf("failed to get all applications: %v", err)
		}
		apps = appList
		return err
	})

	return apps, err

}

func (c *ArgoCDClient) CreateApplication(newApp *application.ApplicationCreateRequest) (*v1alpha1.Application, error) {

	if newApp == nil {
		return nil, errors.New("application must be defined")
	}

	var applicationCreated *v1alpha1.Application
	err := c.DoRequestWithRetry(func(appClient application.ApplicationServiceClient) error {
		app, err := appClient.Create(c.ctx, newApp)
		if err != nil {
			return fmt.Errorf("application can not be created: %v", err)
		}
		applicationCreated = app
		return err
	})

	return applicationCreated, err
}

func (c *ArgoCDClient) GetApplication(query application.ApplicationQuery) (*v1alpha1.Application, error) {

	if isEmpty(query) {
		return nil, errors.New("application name parameter must be defined")
	}

	var foundApp *v1alpha1.Application
	err := c.DoRequestWithRetry(func(appClient application.ApplicationServiceClient) error {
		app, err := appClient.Get(c.ctx, &query)
		if err != nil {
			log.Fatalf("Application not found with query: %v", query)
		}
		foundApp = app
		return nil
	})

	return foundApp, err
}

func isEmpty(query application.ApplicationQuery) bool {
	return query.Name == nil
}

func isUnauthorized(err error) bool {
	return true
}
