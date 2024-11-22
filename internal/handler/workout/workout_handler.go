package workout

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/service"
)

const (
	AddExerciseFailed  = 400
	AddExerciseSuccess = 200
)

type WorkoutHandler struct {
	wrokoutService service.WorkoutService
}

func NewWorkoutHandler(workoutService service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		wrokoutService: workoutService,
	}
}

// @Summary Create empty workout
// @Description Create empty workout
// @Tags workout
// @Accept json
// @Produce json
// @Param request body db.Workout true "Worokout info"
// @Success 200 {object} map[string]interface{} "workout"
// @Router /workout [post]
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
	c.JSON(AddWorkoutSuccess, req)
}

// @Summary Add exercise to workout
// @Description Add exercise to workout
// @Tags workout
// @Accept json
// @Produce json
// @Param workoutID path int true "Workout ID"
// @Param request body db.Exercise true "Exercise info"
// @Success 200 {object} map[string]interface{} "exercise"
// @Router /workout/{workoutID}/exercise [post]
func (w *WorkoutHandler) AddExercise(c *gin.Context) {
	// context.Get("username")
	// c.param("workoutID") uint
	workoutIDParam := c.Param("workoutID")
	workoutID, er := strconv.Atoi(workoutIDParam)
	if er != nil {
		c.JSON(AddExerciseFailed, gin.H{})
		return
	}
	username := c.GetString("username")

	var exercise db.Exercise

	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(AddExerciseFailed, gin.H{})
		return
	}

	er = w.wrokoutService.AddExercise(username, uint(workoutID), &exercise)
	if er != nil {
		c.JSON(AddExerciseFailed, gin.H{})
		return
	}
	c.JSON(AddExerciseSuccess, exercise)
}
