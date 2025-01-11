package service

import (
	"github.com/killerrekt/Go-Fiber-Auth/db"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/request"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/response"
	"github.com/killerrekt/Go-Fiber-Auth/internal/model"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
	"golang.org/x/crypto/bcrypt"
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

func LogIn(req request.LogIn) (response.Standard, error) {
	var user model.User
	err := db.DB.Model(&model.User{}).Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return response.Standard{Message: "Failed to find the user", Status: false}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return response.Standard{Message: "Password don't match", Status: false}, err
	}

	token, err := utils.GenerateAccessToken(user.Email)
	if err != nil {
		return response.Standard{Message: "Failed to create the access token", Status: false}, err
	}

	return response.Standard{Message: "Successfully Logged In", Status: true, Data: token}, nil
}

func GetUser(email string) (model.User, error) {
	var user model.User
	err := db.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return user, err
}
