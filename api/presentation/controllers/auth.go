package controllers

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/requests"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/adapters"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
)

type AuthController struct {
	Jwt            adapters.IJwtAdapter
	AuthRepository interfaces.IRedisAuthRepository
	Service        interfaces.IAuthService
}

func NewAuthController(
	j adapters.IJwtAdapter,
	ra interfaces.IRedisAuthRepository,
	u interfaces.IAuthService) *AuthController {

	return &AuthController{Jwt: j, AuthRepository: ra, Service: u}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	request := new(requests.Login)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.UnprocessableEntity(err))
	}

	err := request.IsValid()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadRequest(err))
	}

	response := ac.Service.Login(request)
	return c.Status(response.Meta.Code).JSON(response)
}

func (ac *AuthController) Logout(c *fiber.Ctx) error {
	token := ac.Jwt.ExtractToken(c)
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized(customErrors.EmptyTokenMessage))
	}

	metadata, err := ac.Jwt.ExtractTokenMetadata(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized(customErrors.InvalidTokenOrExpiredJWTMessage))
	}

	err = ac.AuthRepository.DeleteTokens(c.Context(), metadata)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.InternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Ok("Successfully Logout"))
}

func (ac *AuthController) Refresh(c *fiber.Ctx) error {
	mapToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := c.BodyParser(&mapToken); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.UnprocessableEntity(err))
	}

	refreshToken := mapToken.RefreshToken

	//any error may be due to token expiration
	token, err := ac.Jwt.VerifyRefresh(refreshToken)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized(customErrors.EmptyTokenMessage))
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		//is token valid?
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized(customErrors.EmptyTokenMessage))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.UnprocessableEntity("Unprocessable Entity"))
		}

		userId := claims["user_id"].(string)
		email := claims["email"].(string)

		delErr := ac.AuthRepository.DeleteRefresh(c.Context(), refreshUuid)
		if delErr != nil {
			//Since token is valid, get the uuid:
			//Delete the previous Refresh Token
			return c.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized(customErrors.EmptyTokenMessage))
		}

		ts, createErr := ac.Jwt.CreateTokenJWT(userId, email)
		if createErr != nil {
			return c.Status(fiber.StatusForbidden).JSON(responses.Forbidden("Forbidden"))
		}

		saveErr := ac.AuthRepository.CreateAuth(c.Context(), userId, ts)
		if saveErr != nil {
			return c.Status(fiber.StatusForbidden).JSON(responses.Forbidden("Forbidden"))
		}

		tokens := map[string]string{
			//save the tokens metadata to redis
			//Create new pairs of refresh and access tokens
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		return c.Status(fiber.StatusCreated).JSON(responses.Created(tokens))
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized(customErrors.InvalidTokenOrExpiredJWTMessage))
	}
}
