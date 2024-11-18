package service

import "github.com/kaitokid2302/Workout-Tracker/internal/db"

type Authen struct{}

func (authen *Authen) Login(jwt string) bool {
	return false
}

func (authen *Authen) Register(user db.User) string {
	return ""
}
