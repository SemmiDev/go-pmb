package middlewares

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(environments.AccessSecret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": customErrors.MissingOrMalformedJWTMessage})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": customErrors.InvalidTokenOrExpiredJWTMessage})
}
