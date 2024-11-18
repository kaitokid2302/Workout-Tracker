package service

import (
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
)

type Workout interface {
	CreateWorkout(user string, exercise []*db.Exercise) error
	DeleteExercise(user string, exerciseID uint) error
	AddExercise(user string, exercise *db.Exercise) error
	EditComment(user string, comment string) error
}
