// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/argocdclient/client.go
//
// Generated by this command:
//
//	mockgen -source=pkg/argocdclient/client.go -destination=pkg/argocdclient/client_mock.go -package=argocdclient
//

// Package argocdclient is a generated GoMock package.
package argocdclient

import (
	reflect "reflect"

	application "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	gomock "go.uber.org/mock/gomock"
)

// MockArgoCDClient is a mock of ArgoCDClient interface.
type MockArgoCDClient struct {
	ctrl     *gomock.Controller
	recorder *MockArgoCDClientMockRecorder
	isgomock struct{}
}

// MockArgoCDClientMockRecorder is the mock recorder for MockArgoCDClient.
type MockArgoCDClientMockRecorder struct {
	mock *MockArgoCDClient
}

// NewMockArgoCDClient creates a new mock instance.
func NewMockArgoCDClient(ctrl *gomock.Controller) *MockArgoCDClient {
	mock := &MockArgoCDClient{ctrl: ctrl}
	mock.recorder = &MockArgoCDClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArgoCDClient) EXPECT() *MockArgoCDClientMockRecorder {
	return m.recorder
}

// CreateApplication mocks base method.
func (m *MockArgoCDClient) CreateApplication(newApp *application.ApplicationCreateRequest) (*v1alpha1.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", newApp)
	ret0, _ := ret[0].(*v1alpha1.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockArgoCDClientMockRecorder) CreateApplication(newApp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockArgoCDClient)(nil).CreateApplication), newApp)
}

// DoRequestWithRetry mocks base method.
func (m *MockArgoCDClient) DoRequestWithRetry(requestFunc func(application.ApplicationServiceClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRequestWithRetry", requestFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoRequestWithRetry indicates an expected call of DoRequestWithRetry.
func (mr *MockArgoCDClientMockRecorder) DoRequestWithRetry(requestFunc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRequestWithRetry", reflect.TypeOf((*MockArgoCDClient)(nil).DoRequestWithRetry), requestFunc)
}

// GetApplication mocks base method.
func (m *MockArgoCDClient) GetApplication(query application.ApplicationQuery) (*v1alpha1.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", query)
	ret0, _ := ret[0].(*v1alpha1.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockArgoCDClientMockRecorder) GetApplication(query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockArgoCDClient)(nil).GetApplication), query)
}

// ListApplications mocks base method.
func (m *MockArgoCDClient) ListApplications() (*v1alpha1.ApplicationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications")
	ret0, _ := ret[0].(*v1alpha1.ApplicationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockArgoCDClientMockRecorder) ListApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockArgoCDClient)(nil).ListApplications))
}