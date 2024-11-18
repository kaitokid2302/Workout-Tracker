package repository

import (
	"sort"
	"time"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"gorm.io/gorm"
)

//go:generate mockgen -package=$GOPACKAGE -destination=mock_$GOFILE.go -source=$GOFILE
type WorkoutRepository interface {
	FindUserByUsername(username string) (*db.User, error)
	CreateEmptyWorkout(user *db.User, workoutName string) error
	AddExerciseToWorkout(workout *db.Workout, exercise *db.Exercise) error
	DeleteExercise(workout *db.Workout, exerciseID uint) error
	EditComment(workout *db.Workout, comment string) error
	DeleteWorkout(workout *db.Workout) error
	ScheduleWorkout(workout *db.Workout, time time.Time) error
	FinishWorkout(workout *db.Workout) error
	WorkoutToday(user *db.User) ([]*db.Workout, error)
	WorkoutFuture(user *db.User) ([]*db.Workout, error)
	WorkoutDone(user *db.User) ([]*db.Workout, error)
	PastWorkout(user *db.User) ([]*db.Workout, error)
}

type WorkoutRepositoryImpl struct {
	db *gorm.DB
}

func (w *WorkoutRepositoryImpl) FindUserByUsername(username string) (*db.User, error) {
	client := w.db
	var user db.User
	// res

	er := client.Where("username = ?", username).First(&user).Error
	return &user, er

}

func (w *WorkoutRepositoryImpl) CreateEmptyWorkout(user *db.User, workoutName string) error {
	client := w.db
	workout := db.Workout{
		Name: workoutName,
		// date 2055, moscowtime
		Date:     time.Date(2055, time.January, 1, 0, 0, 0, 0, time.FixedZone("MSK", 3*60*60)),
		Finish:   false,
		Comment:  "",
		UserID:   user.ID,
		Exercise: []db.Exercise{},
	}
	er := client.Save(&workout).Error
	return er
}

func (w *WorkoutRepositoryImpl) AddExerciseToWorkout(workout *db.Workout, exercise *db.Exercise) error {
	client := w.db
	id := workout.ID
	e := db.Exercise{
		Name:       exercise.Name,
		Repetition: exercise.Repetition,
		Set:        exercise.Set,
		Weight:     exercise.Weight,
		WorkoutID:  id,
	}
	er := client.Save(&e).Error
	return er
}

func (w *WorkoutRepositoryImpl) DeleteExercise(workout *db.Workout, exerciseID uint) error {
	client := w.db
	er := client.Where("id = ? AND workout_id = ?", exerciseID, workout.ID).Delete(&db.Exercise{}).Error
	return er
}

func (w *WorkoutRepositoryImpl) EditComment(workout *db.Workout, comment string) error {
	client := w.db
	er := client.Model(workout).Update("comment", comment).Error
	return er
}

func (w *WorkoutRepositoryImpl) DeleteWorkout(workout *db.Workout) error {
	client := w.db
	er := client.Delete(workout).Error
	return er
}

func (w *WorkoutRepositoryImpl) ScheduleWorkout(workout *db.Workout, time time.Time) error {
	client := w.db
	er := client.Model(workout).Update("date", time).Error
	return er
}

func (w *WorkoutRepositoryImpl) FinishWorkout(workout *db.Workout) error {
	client := w.db
	er := client.Model(workout).Update("finish", true).Error
	return er
}

func (w *WorkoutRepositoryImpl) WorkoutToday(user *db.User) ([]*db.Workout, error) {
	client := w.db
	var workout []*db.Workout
	er := client.Where("user_id = ? AND date = ?", user.ID, time.Now().Format("2006-01-02")).Find(&workout).Error
	// sort by time
	sort.Slice(workout, func(i, j int) bool {
		return workout[i].Date.Before(workout[j].Date)
	})
	return workout, er
}

func (w *WorkoutRepositoryImpl) WorkoutFuture(user *db.User) ([]*db.Workout, error) {
	client := w.db
	var workout []*db.Workout
	er := client.Where("user_id = ? AND date > ?", user.ID, time.Now().Format("2006-01-02")).Find(&workout).Error
	sort.Slice(workout, func(i, j int) bool {
		return workout[i].Date.Before(workout[j].Date)
	})
	return workout, er
}

func (w *WorkoutRepositoryImpl) WorkoutDone(user *db.User) ([]*db.Workout, error) {
	client := w.db
	var workout []*db.Workout
	er := client.Where("user_id = ? AND finish = ?", user.ID, true).Find(&workout).Error
	return workout, er
}

func (w *WorkoutRepositoryImpl) PastWorkout(user *db.User) ([]*db.Workout, error) {
	client := w.db
	var workout []*db.Workout
	er := client.Where("user_id = ? AND date < ?", user.ID, time.Now().Format("2006-01-02")).Find(&workout).Error
	return workout, er

}

func NewWorkoutRepository(db *gorm.DB) WorkoutRepository {
	return &WorkoutRepositoryImpl{db: db}
}
