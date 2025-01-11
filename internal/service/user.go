package service

import (
	"github.com/killerrekt/Go-Fiber-Auth/db"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/request"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/response"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
)

func SignUp(req request.SignUp) (response.Standard, error) {
	user := request.SignUpToUser(req)
	err := db.DB.Create(&user).Error
	if err != nil {
		return response.Standard{Message: "Failed to create the user", Status: false}, err
	}

	token, err := utils.GenerateAccessToken(user.Email)
	if err != nil {
		return response.Standard{Message: "Failed to create the access token", Status: false}, err
	}

	return response.Standard{Message: "Successfully Signed Up", Status: true, Data: token}, nil
}
