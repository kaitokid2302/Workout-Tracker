package service

import (
	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
)

type AuthenService struct {
	user *repository.UserService
}

func (authen *AuthenService) Login(jwt string) bool {
	return false
}

func (authen *AuthenService) Register(user db.User) string {
	return ""
}

func NewAuthen(user *repository.UserService) *AuthenService {
	return &AuthenService{user: user}
}
