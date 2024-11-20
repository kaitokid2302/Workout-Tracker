package service

import (
	"errors"
	"testing"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := repository.NewMockUserRepository(ctrl)

	// Case 1: lam - success
	// First check if user exists
	m.EXPECT().FindUser("lam").Return(nil, errors.New("user not found"))
	// Then add user
	m.EXPECT().AddUser(&db.User{Username: "lam"}).Return(nil)

	// Case 2: linh - fail because user exists
	m.EXPECT().FindUser("linh").Return(nil, nil) // user found
	// AddUser won't be called in this case because user exists

	authenService := NewAuthen(m)
	token1, er1 := authenService.Register(db.User{Username: "lam"})
	token2, er2 := authenService.Register(db.User{Username: "linh"})

	assert.Nil(t, er1)
	assert.Error(t, er2)

	assert.NotEmpty(t, token1)
	assert.Empty(t, token2)
}
