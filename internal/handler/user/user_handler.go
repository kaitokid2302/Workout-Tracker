package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/service"
)

type UserHandler struct {
	userService *service.AuthenService
}

func (h *UserHandler) Register(c *gin.Context) {
	// user.UserRegisterRequest
	// map json to struct

	var req UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(RegisterFailed, gin.H{})
		return
	}
}
