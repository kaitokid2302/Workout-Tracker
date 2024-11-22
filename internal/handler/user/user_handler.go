package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/service"
)

type UserHandler struct {
	userService *service.AuthenService
}

func NewUserHandler(userService *service.AuthenService) *UserHandler {
	return &UserHandler{userService: userService}
}

// @Summary Register new user
// @Description Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param request body db.User true "User registration info"
// @Success 200 {object} map[string]interface{} "token"
// @Router /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	// user.UserRegisterRequest
	// map json to struct

	var req db.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(RegisterFailed, gin.H{})
		return
	}
	token, er := h.userService.Register(req)
	if er != nil {
		c.JSON(RegisterFailed, gin.H{})
		return
	}
	c.JSON(RegisterSuccess, gin.H{
		"token": token,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req db.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(LoginFailed, gin.H{})
		return
	}
	ok, token := h.userService.Login(req.Username, req.Password)
	if !ok {
		c.JSON(LoginFailed, gin.H{})
		return
	}

	c.JSON(LoginSuccess, gin.H{
		"token": token,
	})
}
