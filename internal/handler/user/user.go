package user

// define the request and reply struct for user

// login

type UserLoginRequest struct {
	Username string
	Password string
}

type UserLoginReply struct{}

// register

type UserRegisterRequest struct {
	Name     string
	Username string
	Email    string
	Password string
}

type UserRegisterReply struct{}
