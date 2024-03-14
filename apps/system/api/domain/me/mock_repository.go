// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=mock_repository.go -package=me
//

// Package me is a generated GoMock package.
package me

import (
	context "context"
	reflect "reflect"

	provider "github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	account "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	media "github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"
	user "github.com/ryo034/react-go-template/apps/system/api/domain/user"
	invitation "github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	member "github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	bun "github.com/uptrace/bun"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AcceptInvitation mocks base method.
func (m *MockRepository) AcceptInvitation(ctx context.Context, exec bun.IDB, id invitation.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptInvitation", ctx, exec, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptInvitation indicates an expected call of AcceptInvitation.
func (mr *MockRepositoryMockRecorder) AcceptInvitation(ctx, exec, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptInvitation", reflect.TypeOf((*MockRepository)(nil).AcceptInvitation), ctx, exec, id)
}

// ClearMe mocks base method.
func (m *MockRepository) ClearMe(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMe", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearMe indicates an expected call of ClearMe.
func (mr *MockRepositoryMockRecorder) ClearMe(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMe", reflect.TypeOf((*MockRepository)(nil).ClearMe), ctx)
}

// Find mocks base method.
func (m *MockRepository) Find(ctx context.Context, exec bun.IDB, mID member.ID) (*Me, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, exec, mID)
	ret0, _ := ret[0].(*Me)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(ctx, exec, mID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), ctx, exec, mID)
}

// FindAllActiveReceivedInvitations mocks base method.
func (m *MockRepository) FindAllActiveReceivedInvitations(ctx context.Context, exec bun.IDB, aID account.ID) (ReceivedInvitations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllActiveReceivedInvitations", ctx, exec, aID)
	ret0, _ := ret[0].(ReceivedInvitations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllActiveReceivedInvitations indicates an expected call of FindAllActiveReceivedInvitations.
func (mr *MockRepositoryMockRecorder) FindAllActiveReceivedInvitations(ctx, exec, aID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllActiveReceivedInvitations", reflect.TypeOf((*MockRepository)(nil).FindAllActiveReceivedInvitations), ctx, exec, aID)
}

// FindBeforeOnboard mocks base method.
func (m *MockRepository) FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBeforeOnboard", ctx, exec, aID)
	ret0, _ := ret[0].(*Me)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBeforeOnboard indicates an expected call of FindBeforeOnboard.
func (mr *MockRepositoryMockRecorder) FindBeforeOnboard(ctx, exec, aID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBeforeOnboard", reflect.TypeOf((*MockRepository)(nil).FindBeforeOnboard), ctx, exec, aID)
}

// FindByEmail mocks base method.
func (m *MockRepository) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*Me, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, exec, email)
	ret0, _ := ret[0].(*Me)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockRepositoryMockRecorder) FindByEmail(ctx, exec, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockRepository)(nil).FindByEmail), ctx, exec, email)
}

// FindLastLogin mocks base method.
func (m *MockRepository) FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*Me, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLastLogin", ctx, exec, aID)
	ret0, _ := ret[0].(*Me)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLastLogin indicates an expected call of FindLastLogin.
func (mr *MockRepositoryMockRecorder) FindLastLogin(ctx, exec, aID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLastLogin", reflect.TypeOf((*MockRepository)(nil).FindLastLogin), ctx, exec, aID)
}

// RecordLogin mocks base method.
func (m_2 *MockRepository) RecordLogin(ctx context.Context, exec bun.IDB, m *Me) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecordLogin", ctx, exec, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordLogin indicates an expected call of RecordLogin.
func (mr *MockRepositoryMockRecorder) RecordLogin(ctx, exec, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordLogin", reflect.TypeOf((*MockRepository)(nil).RecordLogin), ctx, exec, m)
}

// RemoveProfilePhoto mocks base method.
func (m_2 *MockRepository) RemoveProfilePhoto(ctx context.Context, exec bun.IDB, m *Me) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RemoveProfilePhoto", ctx, exec, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProfilePhoto indicates an expected call of RemoveProfilePhoto.
func (mr *MockRepositoryMockRecorder) RemoveProfilePhoto(ctx, exec, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProfilePhoto", reflect.TypeOf((*MockRepository)(nil).RemoveProfilePhoto), ctx, exec, m)
}

// SetCurrentProvider mocks base method.
func (m *MockRepository) SetCurrentProvider(ctx context.Context, p *provider.Provider) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrentProvider", ctx, p)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// SetCurrentProvider indicates an expected call of SetCurrentProvider.
func (mr *MockRepositoryMockRecorder) SetCurrentProvider(ctx, p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentProvider", reflect.TypeOf((*MockRepository)(nil).SetCurrentProvider), ctx, p)
}

// SetMe mocks base method.
func (m_2 *MockRepository) SetMe(ctx context.Context, m *Me) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SetMe", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMe indicates an expected call of SetMe.
func (mr *MockRepositoryMockRecorder) SetMe(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMe", reflect.TypeOf((*MockRepository)(nil).SetMe), ctx, m)
}

// UpdateMemberProfile mocks base method.
func (m_2 *MockRepository) UpdateMemberProfile(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateMemberProfile", ctx, exec, m)
	ret0, _ := ret[0].(*member.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMemberProfile indicates an expected call of UpdateMemberProfile.
func (mr *MockRepositoryMockRecorder) UpdateMemberProfile(ctx, exec, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMemberProfile", reflect.TypeOf((*MockRepository)(nil).UpdateMemberProfile), ctx, exec, m)
}

// UpdateName mocks base method.
func (m *MockRepository) UpdateName(ctx context.Context, exec bun.IDB, usr *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateName", ctx, exec, usr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateName indicates an expected call of UpdateName.
func (mr *MockRepositoryMockRecorder) UpdateName(ctx, exec, usr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockRepository)(nil).UpdateName), ctx, exec, usr)
}

// UpdateProfilePhoto mocks base method.
func (m_2 *MockRepository) UpdateProfilePhoto(ctx context.Context, exec bun.IDB, m *Me, photo *media.UploadPhoto) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateProfilePhoto", ctx, exec, m, photo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfilePhoto indicates an expected call of UpdateProfilePhoto.
func (mr *MockRepositoryMockRecorder) UpdateProfilePhoto(ctx, exec, m, photo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfilePhoto", reflect.TypeOf((*MockRepository)(nil).UpdateProfilePhoto), ctx, exec, m, photo)
}
