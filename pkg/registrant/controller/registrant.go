package controller

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/pkg/common/config"
	"github.com/SemmiDev/go-pmb/pkg/common/token"
	"github.com/SemmiDev/go-pmb/pkg/common/web"
	"github.com/SemmiDev/go-pmb/pkg/registrant/entity"
	"github.com/SemmiDev/go-pmb/pkg/registrant/models"
	"github.com/SemmiDev/go-pmb/pkg/registrant/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type Controller struct {
	service service.Service
	token   token.Maker
	router  fiber.Router
}

func (s *Controller) Mount(router fiber.Router) {
	router.Post("/registrant/register", s.RegisterHandler)
	router.Post("/registrant/login", s.LoginHandler)
	router.Put("/registrant/payment_status", s.UpdatePaymentStatusHandler)

	s.router = router
}

func NewController(service *service.Service) *Controller {
	// using paseto for token maker.
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalln(fmt.Errorf("cannot create token maker: %w", err))
	}

	return &Controller{
		service: *service,
		token:   tokenMaker,
	}
}

func (s *Controller) RegisterHandler(c *fiber.Ctx) error {
	request := new(models.RegisterReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err.Error()))
	}

	result, err := s.service.Register(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(web.CreatedResponse(result))
}

func (s *Controller) UpdatePaymentStatusHandler(c *fiber.Ctx) error {
	request := new(models.UpdatePaymentStatusReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err.Error()))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err.Error()))
	}

	err = s.service.UpdatePaymentStatus(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(web.OkResponse(nil))
}

func (s *Controller) LoginHandler(c *fiber.Ctx) error {
	request := new(models.LoginReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err.Error()))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err.Error()))
	}

	req := s.service.Repo.FindByUsernameAndPassword(request.Username, request.Password)
	if req.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err.Error()))
	}

	if req.ReadResult.Password == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFoundResponse("registrant not found"))
	}

	if req.ReadResult.PaymentStatus != entity.PaymentStatusPaid {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFoundResponse("please pay the bill first"))
	}

	accessToken, err := s.token.CreateToken(
		req.ReadResult.Username,
		config.AccessTokenDuration,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.NotFoundResponse(err.Error()))
	}

	rsp := models.LoginResponse{
		AccessToken: accessToken,
		Registrant:  req.ReadResult,
	}

	return c.Status(fiber.StatusOK).JSON(web.OkResponse(rsp))
}
