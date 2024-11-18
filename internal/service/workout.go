package service

import (
	"time"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
)

type WorkoutService interface {
	CreateWorkout(user string, exercise []*db.Exercise) error
	DeleteExercise(user string, exerciseID uint) error
	AddExercise(user string, exercise *db.Exercise) error
	EditComment(user string, comment string) error
	DeleteWorkout(user string, exerciseID uint) error
	ScheDuelWorkout(user uint, workoutID uint, time time.Time)
	FinishWorkout(user uint, workoutID uint) error
	WorkoutToday(user uint) []*db.Workout
	WorkoutFuture(user uint) []*db.Workout
	WorkoutDone(user uint) []*db.Workout
	PastWorkout(user uint) []*db.Workout
}
