package workout

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/service"
)

type WorkoutHandler struct {
	wrokoutService service.WorkoutService
}

func NewWorkoutHandler(workoutService service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		wrokoutService: workoutService,
	}
}

func (w *WorkoutHandler) AddWorkout(c *gin.Context) {
	// context.Get("username")
	username := c.GetString("username")
	// create empty workout

	// map json to struct db.Workout

	var req db.Workout
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(AddWorkoutFailed, gin.H{})
		return
	}

	er := w.wrokoutService.CreateWorkout(username, &req)
	if er != nil {
		c.JSON(AddWorkoutFailed, gin.H{})
		return
	}
	c.JSON(AddWorkoutSuccess, gin.H{})
}
