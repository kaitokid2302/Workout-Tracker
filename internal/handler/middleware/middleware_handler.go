package middleware

import "github.com/gin-gonic/gin"

type MiddlewareHandler struct{}

func NewMiddlewareHandler() *MiddlewareHandler {
	return &MiddlewareHandler{}
}

func (m *MiddlewareHandler) SampleAuth(c *gin.Context) {
	c.Status(200)
}
