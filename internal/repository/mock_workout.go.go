// Code generated by MockGen. DO NOT EDIT.
// Source: workout.go
//
// Generated by this command:
//
//	mockgen -package=repository -destination=mock_workout.go.go -source=workout.go
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"
	time "time"

	db "github.com/kaitokid2302/Workout-Tracker/internal/db"
	gomock "go.uber.org/mock/gomock"
)

// MockWorkoutRepository is a mock of WorkoutRepository interface.
type MockWorkoutRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWorkoutRepositoryMockRecorder
}

// MockWorkoutRepositoryMockRecorder is the mock recorder for MockWorkoutRepository.
type MockWorkoutRepositoryMockRecorder struct {
	mock *MockWorkoutRepository
}

// NewMockWorkoutRepository creates a new mock instance.
func NewMockWorkoutRepository(ctrl *gomock.Controller) *MockWorkoutRepository {
	mock := &MockWorkoutRepository{ctrl: ctrl}
	mock.recorder = &MockWorkoutRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkoutRepository) EXPECT() *MockWorkoutRepositoryMockRecorder {
	return m.recorder
}

// AddExerciseToWorkout mocks base method.
func (m *MockWorkoutRepository) AddExerciseToWorkout(workout *db.Workout, exercise *db.Exercise) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddExerciseToWorkout", workout, exercise)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddExerciseToWorkout indicates an expected call of AddExerciseToWorkout.
func (mr *MockWorkoutRepositoryMockRecorder) AddExerciseToWorkout(workout, exercise any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExerciseToWorkout", reflect.TypeOf((*MockWorkoutRepository)(nil).AddExerciseToWorkout), workout, exercise)
}

// CreateEmptyWorkout mocks base method.
func (m *MockWorkoutRepository) CreateEmptyWorkout(user *db.User, workoutName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmptyWorkout", user, workoutName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEmptyWorkout indicates an expected call of CreateEmptyWorkout.
func (mr *MockWorkoutRepositoryMockRecorder) CreateEmptyWorkout(user, workoutName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmptyWorkout", reflect.TypeOf((*MockWorkoutRepository)(nil).CreateEmptyWorkout), user, workoutName)
}

// DeleteExercise mocks base method.
func (m *MockWorkoutRepository) DeleteExercise(workout *db.Workout, exerciseID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExercise", workout, exerciseID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExercise indicates an expected call of DeleteExercise.
func (mr *MockWorkoutRepositoryMockRecorder) DeleteExercise(workout, exerciseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExercise", reflect.TypeOf((*MockWorkoutRepository)(nil).DeleteExercise), workout, exerciseID)
}

// DeleteWorkout mocks base method.
func (m *MockWorkoutRepository) DeleteWorkout(workout *db.Workout) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkout", workout)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkout indicates an expected call of DeleteWorkout.
func (mr *MockWorkoutRepositoryMockRecorder) DeleteWorkout(workout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkout", reflect.TypeOf((*MockWorkoutRepository)(nil).DeleteWorkout), workout)
}

// EditComment mocks base method.
func (m *MockWorkoutRepository) EditComment(workout *db.Workout, comment string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditComment", workout, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditComment indicates an expected call of EditComment.
func (mr *MockWorkoutRepositoryMockRecorder) EditComment(workout, comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditComment", reflect.TypeOf((*MockWorkoutRepository)(nil).EditComment), workout, comment)
}

// FindUserByUsername mocks base method.
func (m *MockWorkoutRepository) FindUserByUsername(username string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUsername", username)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUsername indicates an expected call of FindUserByUsername.
func (mr *MockWorkoutRepositoryMockRecorder) FindUserByUsername(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsername", reflect.TypeOf((*MockWorkoutRepository)(nil).FindUserByUsername), username)
}

// FinishWorkout mocks base method.
func (m *MockWorkoutRepository) FinishWorkout(workout *db.Workout) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishWorkout", workout)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishWorkout indicates an expected call of FinishWorkout.
func (mr *MockWorkoutRepositoryMockRecorder) FinishWorkout(workout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishWorkout", reflect.TypeOf((*MockWorkoutRepository)(nil).FinishWorkout), workout)
}

// PastWorkout mocks base method.
func (m *MockWorkoutRepository) PastWorkout(user *db.User) ([]*db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PastWorkout", user)
	ret0, _ := ret[0].([]*db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PastWorkout indicates an expected call of PastWorkout.
func (mr *MockWorkoutRepositoryMockRecorder) PastWorkout(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PastWorkout", reflect.TypeOf((*MockWorkoutRepository)(nil).PastWorkout), user)
}

// ScheduleWorkout mocks base method.
func (m *MockWorkoutRepository) ScheduleWorkout(workout *db.Workout, time time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleWorkout", workout, time)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleWorkout indicates an expected call of ScheduleWorkout.
func (mr *MockWorkoutRepositoryMockRecorder) ScheduleWorkout(workout, time any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleWorkout", reflect.TypeOf((*MockWorkoutRepository)(nil).ScheduleWorkout), workout, time)
}

// WorkoutDone mocks base method.
func (m *MockWorkoutRepository) WorkoutDone(user *db.User) ([]*db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkoutDone", user)
	ret0, _ := ret[0].([]*db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkoutDone indicates an expected call of WorkoutDone.
func (mr *MockWorkoutRepositoryMockRecorder) WorkoutDone(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkoutDone", reflect.TypeOf((*MockWorkoutRepository)(nil).WorkoutDone), user)
}

// WorkoutFuture mocks base method.
func (m *MockWorkoutRepository) WorkoutFuture(user *db.User) ([]*db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkoutFuture", user)
	ret0, _ := ret[0].([]*db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkoutFuture indicates an expected call of WorkoutFuture.
func (mr *MockWorkoutRepositoryMockRecorder) WorkoutFuture(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkoutFuture", reflect.TypeOf((*MockWorkoutRepository)(nil).WorkoutFuture), user)
}

// WorkoutToday mocks base method.
func (m *MockWorkoutRepository) WorkoutToday(user *db.User) ([]*db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkoutToday", user)
	ret0, _ := ret[0].([]*db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkoutToday indicates an expected call of WorkoutToday.
func (mr *MockWorkoutRepositoryMockRecorder) WorkoutToday(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkoutToday", reflect.TypeOf((*MockWorkoutRepository)(nil).WorkoutToday), user)
}
