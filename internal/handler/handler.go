package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/middleware"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/user"
)

type Hanlder struct {
	userHandler       *user.UserHandler
	middlewareHandler *middleware.MiddlewareHandler
	middleware        *middleware.Middleware
}

func NewHandler(userHandler *user.UserHandler, middlewareHandler *middleware.MiddlewareHandler, middleware *middleware.Middleware) *Hanlder {
	return &Hanlder{userHandler: userHandler, middlewareHandler: middlewareHandler, middleware: middleware}
}

func (h *Hanlder) SetupRoute(engine *gin.Engine) {
	engine.POST("/register", h.userHandler.Register)
	engine.POST("/login", h.userHandler.Login)

	v1 := engine.Group("/v1").Use(h.middleware.Auth())
	{
		v1.GET("/sample", h.middlewareHandler.SampleAuth)
	}
}
