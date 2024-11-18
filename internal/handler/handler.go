package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	App *gin.Engine
}

func (handler *Handler) addRoute() {
	app := handler.App
	app.POST("/workout")

}
