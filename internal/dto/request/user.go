package request

import "github.com/killerrekt/Go-Fiber-Auth/internal/model"

type SignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpToUser(req SignUp) model.User {
	return model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
