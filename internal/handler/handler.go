package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/user"
)

type Hanlder struct {
	userHandler *user.UserHandler
}

func NewHandler(userHandler *user.UserHandler) *Hanlder {
	return &Hanlder{userHandler: userHandler}
}

func (h *Hanlder) SetupRoute(engine *gin.Engine) {
	engine.POST("/register", h.userHandler.Register)
	engine.POST("/login", h.userHandler.Login)
}
