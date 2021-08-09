package controller

import (
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
)

type LoginController struct {
	RegistrationService service.RegistrationService
	Auth                auth.AuthInterface
	Token               auth.TokenInterface
}

func NewAuthController(
	registrationService *service.RegistrationService,
	auth auth.AuthInterface,
	token auth.TokenInterface) LoginController {

	return LoginController{
		RegistrationService: *registrationService,
		Auth:                auth,
		Token:               token,
	}
}

func (c *LoginController) Route(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Post("/auth/login", c.Login)
	v1.Post("/auth/logout", c.Logout)
	v1.Post("/auth/refresh", c.Refresh)
}

func (c *LoginController) Login(ctx *fiber.Ctx) error {
	var request model.LoginRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(model.APIResponse("Cannot unmarshal body", fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	errs := request.Validate()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse(errs, fiber.StatusBadRequest, "Bad Request", nil))
	}

	user, err := c.RegistrationService.Login(&request)
	if err != nil || user == nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse(err.Error(), fiber.StatusBadRequest, "Bad Request", nil))
	}

	ts, tErr := c.Token.CreateToken(user.ID)
	if tErr != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(model.APIResponse(tErr.Error(), fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	saveErr := c.Auth.CreateAuth(ctx.Context(), user.ID, ts)
	if saveErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(model.APIResponse(saveErr.Error(), fiber.StatusInternalServerError, "Internal Server Error", nil))
	}

	response := model.NewLoginResponse(user.ID, user.Name, user.Username, ts.AccessToken, ts.RefreshToken)
	return ctx.Status(fiber.StatusOK).
		JSON(model.APIResponse(nil, fiber.StatusOK, "Ok", response))
}

func (c *LoginController) Logout(ctx *fiber.Ctx) error {
	// check is the user is authenticated first
	metadata, err := c.Token.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(model.APIResponse(err.Error(), fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	//if the access token exist and it is still valid, then delete both the access token and the refresh token
	deleteErr := c.Auth.DeleteTokens(ctx.Context(), metadata)
	if deleteErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(model.APIResponse(deleteErr.Error(), fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(model.APIResponse(nil, fiber.StatusOK, "Ok", map[string]string{
			"message": "Successfully logged out",
		}))
}

//Refresh is the function that uses the refresh_token to generate new pairs of refresh and access tokens.
func (c *LoginController) Refresh(ctx *fiber.Ctx) error {
	mapToken := map[string]string{}

	err := ctx.BodyParser(&mapToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(model.APIResponse(err.Error(), fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	refreshToken := mapToken["refresh_token"]
	if refreshToken == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse("please input the refresh token", fiber.StatusBadRequest, "Bad Request", nil))
	}

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	//any error may be due to token expiration
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(model.APIResponse(err.Error(), fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(model.APIResponse(err.Error(), fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(model.APIResponse("Cannot get uuid", fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
		}

		userId := claims["user_id"].(string)

		//Delete the previous Refresh Token
		delErr := c.Auth.DeleteRefresh(ctx.Context(), refreshUuid)
		if delErr != nil { //if any goes wrong
			return ctx.Status(fiber.StatusUnauthorized).
				JSON(model.APIResponse(delErr.Error(), fiber.StatusUnauthorized, "Unauthorized", nil))
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := c.Token.CreateToken(userId)
		if createErr != nil {
			return ctx.Status(fiber.StatusForbidden).
				JSON(model.APIResponse(createErr.Error(), fiber.StatusForbidden, "Forbidden", nil))
		}

		//save the tokens metadata to redis
		saveErr := c.Auth.CreateAuth(ctx.Context(), userId, ts)
		if saveErr != nil {
			return ctx.Status(fiber.StatusForbidden).
				JSON(model.APIResponse(saveErr.Error(), fiber.StatusForbidden, "Forbidden", nil))
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		return ctx.Status(fiber.StatusCreated).
			JSON(model.APIResponse(nil, fiber.StatusCreated, "Created", tokens))
	} else {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(model.APIResponse("Refresh token expired", fiber.StatusUnauthorized, "Unauthorized", nil))
	}
}
