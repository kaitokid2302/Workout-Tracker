package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
)

type AuthenService struct {
	user repository.UserRepository
}

func (authen *AuthenService) Login(username, password string) (bool, string) {

	userRepository := authen.user

	user, er := userRepository.FindUser(username)
	if er == nil && user.Password == password {
		// create token
		var key = "lam dep trai qua nhi, ahihi, lam yeu lam, lam tung an trom tien cua me"
		var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss":  "mra2322001",
			"sub":  user.Username,
			"role": "user",
			"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		s, _ := t.SignedString([]byte(key))
		return true, s
	}
	return false, ""

}

func (authen *AuthenService) Register(user db.User) (string, error) {
	if _, er := authen.user.FindUser(user.Username); er != nil {
		var key = "lam dep trai qua nhi, ahihi, lam yeu lam, lam tung an trom tien cua me"
		var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss":  "mra2322001",
			"sub":  user.Username,
			"role": "user",
			"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		s, _ := t.SignedString([]byte(key))
		authen.user.AddUser(&user)
		return s, nil
	}
	return "", errors.New("username existed")
}

func NewAuthen(user repository.UserRepository) *AuthenService {
	return &AuthenService{user: user}
}
