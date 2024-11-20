// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -package=repository -destination=mock_user.go.go -source=user.go
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	db "github.com/kaitokid2302/Workout-Tracker/internal/db"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserRepository) AddUser(user *db.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserRepositoryMockRecorder) AddUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserRepository)(nil).AddUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUser")
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser))
}

// FindUser mocks base method.
func (m *MockUserRepository) FindUser(username string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", username)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserRepositoryMockRecorder) FindUser(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserRepository)(nil).FindUser), username)
}
