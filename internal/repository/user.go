package repository

import (
	"errors"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"gorm.io/gorm"
)

//go:generate mockgen -package=$GOPACKAGE -destination=mock_$GOFILE.go -source=$GOFILE
type UserRepository interface {
	AddUser(user *db.User) error
	FindUser(username string) (*db.User, error)
	DeleteUser()
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) AddUser(user *db.User) error {
	var puser db.User
	u.db.Where("user_name = ?", user.Username).First(&puser)
	if puser.Username != "" {
		return errors.New("username exists")
	}
	u.db.Save(user)
	return nil
}

func (u *UserRepositoryImpl) FindUser(username string) (*db.User, error) {
	var user db.User
	u.db.Where("user_name = ?", username).First(&user)
	if user.Username == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
