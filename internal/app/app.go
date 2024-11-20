package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/user"
)

type App struct {
	server      *gin.Engine
	userHandler *user.UserHandler
}

func (app *App) SetupRoute() {
	app.server = gin.Default()

	// re
}
