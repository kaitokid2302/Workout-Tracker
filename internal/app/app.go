package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler"
	"github.com/kaitokid2302/Workout-Tracker/internal/handler/user"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
	"github.com/kaitokid2302/Workout-Tracker/internal/service"
)

type App struct {
	server  *gin.Engine
	handler handler.Hanlder
}

func (app *App) SetUp() {
	clientDB := db.DBInit()
	app.server = gin.Default()

	userRepository := repository.NewUserRepository(clientDB)
	userService := service.NewAuthen(userRepository)
	userHandler := user.NewUserHandler(userService)
	app.handler = *handler.NewHandler(userHandler)

	app.handler.SetupRoute(app.server)
}

func (app *App) Run() {
	app.server.Run()
}
