package registrant

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/pkg/common/config"
	"github.com/SemmiDev/go-pmb/pkg/common/token"
	"github.com/SemmiDev/go-pmb/pkg/common/web"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type Server struct {
	service Service
	token   token.Maker
	router  fiber.Router
}

func (s *Server) Mount(router fiber.Router) {
	router.Post("/registrant/register", s.RegisterHandler)
	router.Post("/registrant/login", s.LoginHandler)
	router.Put("/registrant/payment_status", s.UpdatePaymentStatusHandler)

	s.router = router
}

func NewServer(service *Service) *Server {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalln(fmt.Errorf("cannot create token maker: %w", err))
	}

	return &Server{
		service: *service,
		token:   tokenMaker,
	}
}

func (s *Server) RegisterHandler(c *fiber.Ctx) error {
	request := new(RegisterReq)
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

func (s *Server) UpdatePaymentStatusHandler(c *fiber.Ctx) error {
	request := new(UpdatePaymentStatusReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err.Error()))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err.Error()))
	}

	req := s.service.repo.FindByID(request.RegisterID)
	if req.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err.Error()))
	}

	if req.ReadResult.Email == "" {
		return c.Status(http.StatusNotFound).JSON(web.NotFoundResponse(err.Error()))
	}

	err = s.service.UpdatePaymentStatus(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(web.OkResponse(nil))
}

func (s *Server) LoginHandler(c *fiber.Ctx) error {
	request := new(LoginReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err.Error()))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err.Error()))
	}

	req := s.service.repo.FindByUsernameAndPassword(request.Username, request.Password)
	if req.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err.Error()))
	}

	if req.ReadResult.Password == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFoundResponse("registrant not found"))
	}

	if req.ReadResult.PaymentStatus != PaymentStatusPaid {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFoundResponse("please pay the bill first"))
	}

	accessToken, err := s.token.CreateToken(
		req.ReadResult.Username,
		config.AccessTokenDuration,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.NotFoundResponse(err.Error()))
	}

	rsp := LoginResponse{
		AccessToken: accessToken,
		Registrant:  req.ReadResult,
	}

	return c.Status(fiber.StatusOK).JSON(web.OkResponse(rsp))
}
