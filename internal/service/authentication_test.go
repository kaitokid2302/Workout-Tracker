package service

import (
	"errors"
	"testing"

	"github.com/kaitokid2302/Workout-Tracker/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
	authen := AuthenService{}
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	assert.True(t, authen.Login(s))

	s = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5d"

	assert.False(t, authen.Login(s))
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := repository.NewMockUserService(ctrl)

	// Expect add user: "lam" to return nil
	// Expect find user: "linh" to return error
	// expect add user: "thoa" to return error

	m.EXPECT().AddUser("lam").Return(nil)
	m.EXPECT().FindUser("linh").Return(errors.New("user not found"))
	m.EXPECT().AddUser("thoa").Return(errors.New("username exists"))

	authenService := NewAuthen(m)
}
