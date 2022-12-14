// Code generated by MockGen. DO NOT EDIT.
// Source: jira.go

// Package mocks is a generated GoMock package.
package mocks

import (
	jira_go_api "gopkg.in/andygrunwald/go-jira.v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockissuer is a mock of issuer interface.
type Mockissuer struct {
	ctrl     *gomock.Controller
	recorder *MockissuerMockRecorder
}

// MockissuerMockRecorder is the mock recorder for Mockissuer.
type MockissuerMockRecorder struct {
	mock *Mockissuer
}

// NewMockissuer creates a new mock instance.
func NewMockissuer(ctrl *gomock.Controller) *Mockissuer {
	mock := &Mockissuer{ctrl: ctrl}
	mock.recorder = &MockissuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockissuer) EXPECT() *MockissuerMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *Mockissuer) Search(jql string, options *jira_go_api.SearchOptions) ([]jira_go_api.Issue, *jira_go_api.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", jql, options)
	ret0, _ := ret[0].([]jira_go_api.Issue)
	ret1, _ := ret[1].(*jira_go_api.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Search indicates an expected call of Search.
func (mr *MockissuerMockRecorder) Search(jql, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*Mockissuer)(nil).Search), jql, options)
}
