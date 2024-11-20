package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/middleware"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/user"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/workout"
)

type Hanlder struct {
	userHandler       *user.UserHandler
	middlewareHandler *middleware.MiddlewareHandler
	middleware        *middleware.Middleware
	workoutHandler    *workout.WorkoutHandler
}

func NewHandler(userHandler *user.UserHandler, middlewareHandler *middleware.MiddlewareHandler, middleware *middleware.Middleware, workoutHandler *workout.WorkoutHandler) *Hanlder {
	return &Hanlder{userHandler: userHandler, middlewareHandler: middlewareHandler, middleware: middleware, workoutHandler: workoutHandler}
}

func (h *Hanlder) SetupRoute(engine *gin.Engine) {
	engine.POST("/register", h.userHandler.Register)
	engine.POST("/login", h.userHandler.Login)

	v1 := engine.Group("/v1").Use(h.middleware.Auth())
	{
		v1.GET("/sample", h.middlewareHandler.SampleAuth)
		v1.POST("/workout", h.workoutHandler.AddWorkout)
		// add exercise to workout by id
		v1.POST("/workout/:id/exercise", h.workoutHandler.AddExercise)
	}
}
