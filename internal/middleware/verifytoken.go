package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/response"
	"github.com/killerrekt/Go-Fiber-Auth/internal/service"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
)

func AuthenticateAndAuthorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		config := utils.Config
		accessSecret := config.AccessTokenSecret

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Standard{Message: "Authorization header is missing", Status: false})
		}

		accessToken := authHeader[len("Bearer "):]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Standard{Message: "Expired Access Token", Status: false})
		}

		email, ok := claims["email"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Standard{Message: "Invalid Access Token", Status: false})
		}

		user, err := service.GetUser(email)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(response.Standard{Message: "Failed to get user from token", Status: false})
		}

		c.Locals("auth-user", user)

		return c.Next()
	}
}
