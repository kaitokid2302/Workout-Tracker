package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	Unauthorized = 401
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get header authorization
		s := c.GetHeader("Authorization")
		key := "lam dep trai qua nhi, ahihi, lam yeu lam, lam tung an trom tien cua me"
		token, er := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("can not decode")
			}
			return []byte(key), nil
		})
		if er != nil {
			c.AbortWithStatus(Unauthorized)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims.Valid() != nil {
				c.AbortWithStatus(Unauthorized)
				return
			}
			username, ok := claims["username"].(string)
			if ok {
				c.Set("username", username)
				c.Next()
			} else {
				c.AbortWithStatus(Unauthorized)
				return
			}
		} else {
			c.AbortWithStatus(Unauthorized)
			return
		}

	}
}
