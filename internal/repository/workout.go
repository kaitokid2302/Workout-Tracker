package repository

import (
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"gorm.io/gorm"
)

type WorkoutRepository interface {
	FindUserByUsername(username string) (*db.User, error)
	CreateEmptyWorkout(user *db.User) error
	AddExerciseToWorkout(workout *db.Workout, exercise *db.Exercise) error
	DeleteExercise(workout *db.Workout, exerciseID uint) error
	EditComment(workout *db.Workout, comment string) error
	DeleteWorkout(workout *db.Workout) error
	ScheduleWorkout(workout *db.Workout, time string) error
	FinishWorkout(workout *db.Workout) error
	WorkoutToday(user *db.User) ([]*db.Workout, error)
	WorkoutFuture(user *db.User) ([]*db.Workout, error)
	WorkoutDone(user *db.User) ([]*db.Workout, error)
	PastWorkout(user *db.User) ([]*db.Workout, error)
}

type WorkoutRepositoryImpl struct {
	db *gorm.DB
}
