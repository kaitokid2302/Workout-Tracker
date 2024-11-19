package service

import (
	"time"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
)

type WorkoutService interface {
	CreateWorkout(username string, exercise []*db.Exercise) error
	DeleteExercise(username string, exerciseID uint) error
	AddExercise(username string, exercise *db.Exercise) error
	EditComment(username string, comment string) error
	DeleteWorkout(username string, exerciseID uint) error
	ScheDuelWorkout(userID uint, workoutID uint, time time.Time)
	FinishWorkout(userOD uint, workoutID uint) error
	WorkoutToday(userID uint) []*db.Workout
	WorkoutFuture(userID uint) []*db.Workout
	WorkoutDone(userID uint) []*db.Workout
	PastWorkout(userID uint) []*db.Workout
}

type WorkoutServiceImpl struct {
	workoutRepository repository.WorkoutRepository
	userRepository    repository.UserRepository
}

func (w *WorkoutServiceImpl) CreateWorkout(username string, exercise []*db.Exercise) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) DeleteExercise(username string, exerciseID uint) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) AddExercise(username string, exercise *db.Exercise) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) EditComment(username string, comment string) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) DeleteWorkout(username string, exerciseID uint) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) ScheDuelWorkout(userID uint, workoutID uint, time time.Time) {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) FinishWorkout(userOD uint, workoutID uint) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) WorkoutToday(userID uint) []*db.Workout {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) WorkoutFuture(userID uint) []*db.Workout {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) WorkoutDone(userID uint) []*db.Workout {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) PastWorkout(userID uint) []*db.Workout {
	panic("not implemented") // TODO: Implement
}
