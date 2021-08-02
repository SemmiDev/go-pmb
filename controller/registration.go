package controller

import (
	"errors"
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
)

type RegistrationController struct {
	RegistrationService model.RegistrationService
	Auth                auth.AuthInterface
	Token               auth.TokenInterface
}

func NewRegistrationController(registrationService *model.RegistrationService, auth auth.AuthInterface, token auth.TokenInterface,
) RegistrationController {
	return RegistrationController{
		RegistrationService: *registrationService,
		Auth:                auth,
		Token:               token,
	}
}

func (c *RegistrationController) Route(app *fiber.App) {

	v1 := app.Group("/api/v1")

	v1.Post("/auth/login", c.Login)
	v1.Post("/auth/logout", c.Logout)
	v1.Post("/auth/refresh", c.Refresh)

	v1.Post("/registration", c.Register)
	v1.Put("/registration/status", c.UpdateStatusBilling)
}

func (c *RegistrationController) Register(ctx *fiber.Ctx) error {
	var request model.RegistrationRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:         fiber.StatusUnprocessableEntity,
			Status:       "Unprocessable Entity",
			Error:        true,
			ErrorMessage: "Cannot unmarshal body",
			Data:         nil,
		})
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: errs,
			Data:         nil,
		})
	}

	var program model.Program
	switch request.Program {
	case "S1D3D4":
		program = model.S1D3D4
	case "S2":
		program = model.S2
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Error:  true,
			ErrorMessage: map[string]string{
				"Program_Not_Available": "Please Chose Between S1D3D4 or S2",
			},
			Data: nil,
		})
	}

	response, err := c.RegistrationService.Create(&request, program)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:         fiber.StatusCreated,
		Status:       "Created",
		Error:        false,
		ErrorMessage: nil,
		Data:         response,
	})
}

func (c *RegistrationController) UpdateStatusBilling(ctx *fiber.Ctx) error {
	var request model.UpdateStatus

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:         fiber.StatusUnprocessableEntity,
			Status:       "Unprocessable Entity",
			Error:        true,
			ErrorMessage: "Cannot unmarshal body",
			Data:         nil,
		})
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: errs,
			Data:         nil,
		})
	}

	err = c.RegistrationService.UpdateStatusBilling(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Code:         fiber.StatusInternalServerError,
			Status:       "Internal Server Error",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:         fiber.StatusOK,
		Status:       "Ok",
		Error:        false,
		ErrorMessage: nil,
		Data: fiber.Map{
			"status": "updated",
		},
	})
}

func (c *RegistrationController) Login(ctx *fiber.Ctx) error {
	var request model.LoginRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:         fiber.StatusUnprocessableEntity,
			Status:       "Unprocessable Entity",
			Error:        true,
			ErrorMessage: "Cannot unmarshal body",
			Data:         nil,
		})
	}

	errs := request.Validate()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: errs,
			Data:         nil,
		})
	}

	user, err := c.RegistrationService.GetByUsername(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	if user.Status == false {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: "please pay the billing first",
			Data:         nil,
		})
	}

	err = util.Check(request.Password, user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	ts, tErr := c.Token.CreateToken(user.ID)
	if tErr != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:         fiber.StatusUnprocessableEntity,
			Status:       "Unprocessable Entity",
			Error:        true,
			ErrorMessage: tErr.Error(),
			Data:         nil,
		})
	}

	saveErr := c.Auth.CreateAuth(ctx.Context(), user.ID, ts)
	if saveErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Code:         fiber.StatusInternalServerError,
			Status:       "Internal Server Error",
			Error:        true,
			ErrorMessage: saveErr.Error(),
			Data:         nil,
		})
	}

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = user.ID
	userData["username"] = user.Username
	userData["name"] = user.Name

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:         fiber.StatusOK,
		Status:       "Ok",
		Error:        false,
		ErrorMessage: nil,
		Data:         userData,
	})
}

func (c *RegistrationController) Logout(ctx *fiber.Ctx) error {
	//check is the user is authenticated first
	metadata, err := c.Token.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
			Code:         fiber.StatusUnauthorized,
			Status:       "Unauthorized",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	//if the access token exist and it is still valid, then delete both the access token and the refresh token
	deleteErr := c.Auth.DeleteTokens(ctx.Context(), metadata)
	if deleteErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
			Code:         fiber.StatusUnauthorized,
			Status:       "Unauthorized",
			Error:        true,
			ErrorMessage: deleteErr.Error(),
			Data:         nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:         fiber.StatusOK,
		Status:       "Ok",
		Error:        false,
		ErrorMessage: nil,
		Data: fiber.Map{
			"message": "Successfully logged out",
		},
	})
}

//Refresh is the function that uses the refresh_token to generate new pairs of refresh and access tokens.
func (c *RegistrationController) Refresh(ctx *fiber.Ctx) error {
	mapToken := map[string]string{}
	err := ctx.BodyParser(&mapToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:         fiber.StatusUnprocessableEntity,
			Status:       "Unprocessable Entity",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	refreshToken := mapToken["refresh_token"]
	if refreshToken == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: errors.New("please input the refresh token").Error(),
			Data:         nil,
		})
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
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
			Code:         fiber.StatusUnauthorized,
			Status:       "Unauthorized",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
			Code:         fiber.StatusUnauthorized,
			Status:       "Unauthorized",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
				Code:         fiber.StatusUnprocessableEntity,
				Status:       "Unprocessable Entity",
				Error:        true,
				ErrorMessage: "Cannot get uuid",
				Data:         nil,
			})
		}

		userId := claims["user_id"].(string)
		//Delete the previous Refresh Token
		delErr := c.Auth.DeleteRefresh(ctx.Context(), refreshUuid)
		if delErr != nil { //if any goes wrong
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
				Code:         fiber.StatusUnauthorized,
				Status:       "Unauthorized",
				Error:        true,
				ErrorMessage: "Unauthorized",
				Data:         nil,
			})
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := c.Token.CreateToken(userId)
		if createErr != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(model.WebResponse{
				Code:         fiber.StatusForbidden,
				Status:       "Forbidden",
				Error:        true,
				ErrorMessage: createErr.Error(),
				Data:         nil,
			})
		}

		//save the tokens metadata to redis
		saveErr := c.Auth.CreateAuth(ctx.Context(), userId, ts)
		if saveErr != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(model.WebResponse{
				Code:         fiber.StatusForbidden,
				Status:       "Forbidden",
				Error:        true,
				ErrorMessage: saveErr.Error(),
				Data:         nil,
			})
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
			Code:         fiber.StatusCreated,
			Status:       "Created",
			Error:        false,
			ErrorMessage: nil,
			Data:         tokens,
		})

	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
			Code:         fiber.StatusUnauthorized,
			Status:       "Unauthorized",
			Error:        true,
			ErrorMessage: "Refresh token expired",
			Data:         nil,
		})
	}
}
