// Code generated by MockGen. DO NOT EDIT.
// Source: jira.go

// Package mocks is a generated GoMock package.
package mocks

import (
	jira_go_api "gopkg.in/andygrunwald/go-jira.v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJiraer is a mock of Jiraer interface.
type MockJiraer struct {
	ctrl     *gomock.Controller
	recorder *MockJiraerMockRecorder
}

// MockJiraerMockRecorder is the mock recorder for MockJiraer.
type MockJiraerMockRecorder struct {
	mock *MockJiraer
}

// NewMockJiraer creates a new mock instance.
func NewMockJiraer(ctrl *gomock.Controller) *MockJiraer {
	mock := &MockJiraer{ctrl: ctrl}
	mock.recorder = &MockJiraerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJiraer) EXPECT() *MockJiraerMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockJiraer) Search(jql string) ([]jira_go_api.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", jql)
	ret0, _ := ret[0].([]jira_go_api.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockJiraerMockRecorder) Search(jql interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockJiraer)(nil).Search), jql)
}