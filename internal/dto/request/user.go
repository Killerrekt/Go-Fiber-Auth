package request

import "github.com/killerrekt/Go-Fiber-Auth/internal/model"

type SignUp struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func SignUpToUser(req SignUp) model.User {
	return model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type LogIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResetPassword struct {
	NewPassword string `json:"new_password" validate:"required"`
}
