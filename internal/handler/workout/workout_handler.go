package workout

import "github.com/kaitokid2302/Workout-Tracker/internal/service"

type WorkoutHandler struct {
	wrokoutService service.WorkoutService
}

func NewWorkoutHandler(workoutService service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		wrokoutService: workoutService,
	}
}
