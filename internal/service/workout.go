package service

import (
	"errors"
	"time"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
)

type WorkoutService interface {
	CreateWorkout(username string, workout *db.Workout) error
	DeleteExercise(username string, exerciseID uint) error
	AddExercise(username string, workoutID uint, exercise *db.Exercise) error
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

func (w *WorkoutServiceImpl) CreateWorkout(username string, workout *db.Workout) error {
	user, err := w.workoutRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	err = w.workoutRepository.CreateEmptyWorkout(user, workout.Name)
	if err != nil {
		return err
	}
	if workout.Exercise != nil {
		for _, exercise := range workout.Exercise {
			err = w.workoutRepository.AddExerciseToWorkout(workout, &exercise)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *WorkoutServiceImpl) DeleteExercise(username string, exerciseID uint) error {
	panic("not implemented") // TODO: Implement
}

func (w *WorkoutServiceImpl) AddExercise(username string, workoutID uint, exercise *db.Exercise) error {
	workout, er := w.workoutRepository.FindWorkoutByID(workoutID)
	if er != nil {
		return er
	}
	er = w.workoutRepository.AddExerciseToWorkout(workout, exercise)
	if er != nil {
		return er
	}
	user, er := w.workoutRepository.FindUserByUsername(username)
	if er != nil {
		return er
	}
	if user.ID != workout.UserID {
		return errors.New("user is not owner of workout")
	}
	return nil
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

func NewWorkoutService(workoutRepository repository.WorkoutRepository, userRepository repository.UserRepository) WorkoutService {
	return &WorkoutServiceImpl{
		workoutRepository: workoutRepository,
		userRepository:    userRepository,
	}
}
