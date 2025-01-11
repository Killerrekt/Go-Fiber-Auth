package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/request"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/response"
	"github.com/killerrekt/Go-Fiber-Auth/internal/model"
	"github.com/killerrekt/Go-Fiber-Auth/internal/service"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var req request.SignUp

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Standard{Message: "Failed to parse the json", Status: false})
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusConflict).JSON(response.Standard{Message: "Missing some fields :-" + err.Error(), Status: false})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusFailedDependency).JSON(response.Standard{Message: "Failed to hash the password", Status: false})
	}

	req.Password = string(hashedPassword)

	res, err := service.SignUp(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusAccepted).JSON(res)
}

func LogIn(c *fiber.Ctx) error {
	var req request.LogIn

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Standard{Message: "Failed to parse the json", Status: false})
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusConflict).JSON(response.Standard{Message: "Missing some fields :-" + err.Error(), Status: false})
	}

	res, err := service.LogIn(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusAccepted).JSON(res)
}

func Me(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(response.Standard{Message: "The user details are as followed", Status: true, Data: c.Locals("auth-user")})
}

func ResetPassword(c *fiber.Ctx) error {
	var req request.ResetPassword

	user := c.Locals("auth-user").(model.User)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Standard{Message: "Failed to parse the json", Status: false})
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusConflict).JSON(response.Standard{Message: "Missing some fields :-" + err.Error(), Status: false})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusFailedDependency).JSON(response.Standard{Message: "Failed to hash the password", Status: false})
	}
	req.NewPassword = string(hashedPassword)

	res, err := service.ResetPassword(user.Email, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusAccepted).JSON(res)
}
