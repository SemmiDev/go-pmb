package registrant

import (
	"errors"
	"github.com/SemmiDev/go-pmb/src/common/config"
	"github.com/SemmiDev/go-pmb/src/common/helper"
	"github.com/SemmiDev/go-pmb/src/common/web"
	"github.com/SemmiDev/go-pmb/src/registrant/entities"
	"github.com/SemmiDev/go-pmb/src/registrant/models"
	"github.com/SemmiDev/go-pmb/src/registrant/payments"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

func (s *Server) HandleRegisterRegistrant(c *fiber.Ctx) error {
	request := new(models.RegisterReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err))
	}

	result, err := s.registerRegistrant(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err))
	}

	return c.Status(http.StatusCreated).JSON(web.CreatedResponse(result))
}

func (s *Server) registerRegistrant(data *models.RegisterReq) (response *models.RegisterResponse, err error) {
	password := helper.GeneratePassword()
	result := entities.NewRegistrant(data.Name, data.Email, data.Phone, password, data.Program)

	paymentPayload := payments.NewPayload(uuid.NewString(), result.Bill.Val())
	paymentURL, err := s.midtrans.GetPaymentURL(paymentPayload, result)
	if err != nil {
		return nil, err
	}
	result.PaymentURL = paymentURL

	err = <-s.registrantRepo.Save(result)
	if err != nil {
		return nil, err
	}

	return models.ToRegisterRegistrantResp(result, password, paymentPayload.AmountFormatIDR()), nil
}

func (s *Server) HandleUpdateRegisterPaymentStatus(c *fiber.Ctx) error {
	request := new(models.UpdatePaymentStatusReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err))
	}

	req := <-s.registrantRepo.GetByID(request.RegisterID)
	if req.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err))
	}

	registrantData := req.Result.(models.ReadResult)
	if registrantData.Email == "" {
		return c.Status(http.StatusNotFound).JSON(web.NotFoundResponse(err))
	}

	var status entities.PaymentStatus
	if request.PaymentType == "credit_card" && request.PaymentStatus == "capture" && request.FraudStatus == "accept" {
		status = entities.PaymentStatusPaid
	} else if request.PaymentStatus == "settlement" {
		status = entities.PaymentStatusPaid
	} else if request.PaymentStatus == "deny" || request.PaymentStatus == "expire" || request.PaymentStatus == "cancel" {
		status = entities.PaymentStatusCancel
	} else {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err))
	}

	err = <-s.registrantRepo.UpdateStatus(request.RegisterID, status)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(web.OkResponse(nil))
}

func (s *Server) HandleRegistrantLogin(c *fiber.Ctx) error {
	request := new(models.LoginReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(web.UnprocessableEntityResponse(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BadRequestResponse(err))
	}

	req := <-s.registrantRepo.GetByUsernameAndPassword(request.Username, request.Password)
	if req.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.ResponseInternalServerError(err))
	}

	registrantData := req.Result.(models.ReadResult)
	if registrantData.Password == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFoundResponse(errors.New("registrant not found")))
	}

	accessToken, err := s.tokenMaker.CreateToken(
		registrantData.Username,
		config.AccessTokenDuration,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(web.NotFoundResponse(err))
	}

	rsp := models.LoginResponse{
		AccessToken: accessToken,
		Registrant:  registrantData,
	}

	return c.Status(fiber.StatusOK).JSON(web.OkResponse(rsp))
}
