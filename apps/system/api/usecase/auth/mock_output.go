// Code generated by MockGen. DO NOT EDIT.
// Source: output.go
//
// Generated by this command:
//
//	mockgen -source=output.go -destination=mock_output.go -package=auth
//

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	me "github.com/ryo034/react-go-template/apps/system/api/domain/me"
	openapi "github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	gomock "go.uber.org/mock/gomock"
)

// MockOutputPort is a mock of OutputPort interface.
type MockOutputPort struct {
	ctrl     *gomock.Controller
	recorder *MockOutputPortMockRecorder
}

// MockOutputPortMockRecorder is the mock recorder for MockOutputPort.
type MockOutputPortMockRecorder struct {
	mock *MockOutputPort
}

// NewMockOutputPort creates a new mock instance.
func NewMockOutputPort(ctrl *gomock.Controller) *MockOutputPort {
	mock := &MockOutputPort{ctrl: ctrl}
	mock.recorder = &MockOutputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutputPort) EXPECT() *MockOutputPortMockRecorder {
	return m.recorder
}

// AuthByAuth mocks base method.
func (m *MockOutputPort) AuthByAuth(me *me.Me) (openapi.APIV1AuthOAuthPostRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthByAuth", me)
	ret0, _ := ret[0].(openapi.APIV1AuthOAuthPostRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthByAuth indicates an expected call of AuthByAuth.
func (mr *MockOutputPortMockRecorder) AuthByAuth(me any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthByAuth", reflect.TypeOf((*MockOutputPort)(nil).AuthByAuth), me)
}

// InvitationByToken mocks base method.
func (m *MockOutputPort) InvitationByToken(ri me.ReceivedInvitation) (openapi.GetInvitationByTokenRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvitationByToken", ri)
	ret0, _ := ret[0].(openapi.GetInvitationByTokenRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvitationByToken indicates an expected call of InvitationByToken.
func (mr *MockOutputPortMockRecorder) InvitationByToken(ri any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvitationByToken", reflect.TypeOf((*MockOutputPort)(nil).InvitationByToken), ri)
}

// JwtToken mocks base method.
func (m *MockOutputPort) JwtToken(token string) *openapi.JwtToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JwtToken", token)
	ret0, _ := ret[0].(*openapi.JwtToken)
	return ret0
}

// JwtToken indicates an expected call of JwtToken.
func (mr *MockOutputPortMockRecorder) JwtToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JwtToken", reflect.TypeOf((*MockOutputPort)(nil).JwtToken), token)
}

// ProcessInvitationOAuth mocks base method.
func (m *MockOutputPort) ProcessInvitationOAuth(me *me.Me) (openapi.ProcessInvitationOAuthRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessInvitationOAuth", me)
	ret0, _ := ret[0].(openapi.ProcessInvitationOAuthRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessInvitationOAuth indicates an expected call of ProcessInvitationOAuth.
func (mr *MockOutputPortMockRecorder) ProcessInvitationOAuth(me any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessInvitationOAuth", reflect.TypeOf((*MockOutputPort)(nil).ProcessInvitationOAuth), me)
}
