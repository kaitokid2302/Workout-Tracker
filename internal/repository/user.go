package repository

import (
	"errors"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"gorm.io/gorm"
)

//go:generate mockgen -package=$GOPACKAGE -destination=mock_$GOFILE.go -source=$GOFILE
type UserService interface {
	AddUser(user *db.User) error
	FindUser(username string) error
}

type UserServiceImpl struct {
	db *gorm.DB
}

func (u *UserServiceImpl) AddUser(user *db.User) error {
	var puser db.User
	u.db.Where("user_name = ?", user.Username).First(&puser)
	if puser.Username != "" {
		return errors.New("username exists")
	}
	u.db.Save(user)
	return nil
}

func (u *UserServiceImpl) FindUser(username string) error {
	var user db.User
	u.db.Where("user_name = ?", username).First(&user)
	if user.Username == "" {
		return errors.New("user not found")
	}
	return nil
}
