package service

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
)

type AuthenService struct {
	user repository.UserService
}

func (authen *AuthenService) Login(s string) bool {
	var key = "lam dep trai qua nhi, ahihi, lam yeu lam, lam tung an trom tien cua me"
	token, er := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("can not decode")
		}
		return []byte(key), nil
	})
	if er != nil {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims.Valid() == nil
	} else {
		return false
	}

}

func (authen *AuthenService) Register(user db.User) (string, error) {
	if authen.user.FindUser(user.Username) != nil {
		var key = "lam dep trai qua nhi, ahihi, lam yeu lam, lam tung an trom tien cua me"
		var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss":  "mra2322001",
			"role": "user",
			"exp":  24 * 3600,
		})
		s, _ := t.SignedString([]byte(key))
		authen.user.AddUser(&user)
		return s, nil
	}
	return "", errors.New("username existed")
}

func NewAuthen(user repository.UserService) *AuthenService {
	return &AuthenService{user: user}
}
