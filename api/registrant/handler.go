package registrant

import (
	"errors"
	"github.com/SemmiDev/go-pmb/common/config"
	"github.com/SemmiDev/go-pmb/common/helper"
	"github.com/SemmiDev/go-pmb/common/web"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) HandleRegisterRegistrant(c *fiber.Ctx) error {
	request := new(RegisterReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(web.UnprocessableEntity(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.BadRequest(err))
	}

	result, err := s.registerRegistrant(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	return c.Status(fiber.StatusCreated).JSON(web.Created(result))
}

func (s *Server) registerRegistrant(data *RegisterReq) (response *RegisterResponse, err error) {
	password := helper.GeneratePassword()
	result := New(data.Name, data.Email, data.Phone, password, data.Program)

	paymentPayload := payload{
		id:     uuid.NewString(),
		amount: result.bill.Value(),
	}

	paymentURL, err := s.midtrans.GetPaymentURL(&paymentPayload, result)
	if err != nil {
		return nil, err
	}
	result.paymentURL = paymentURL

	err = <-s.registrantRepo.Cmd.Save(result)
	if err != nil {
		return nil, err
	}

	return ToRegisterRegistrantResp(result, password, paymentPayload.AmountFormatIDR()), nil
}

func (s *Server) HandleUpdateRegisterPaymentStatus(c *fiber.Ctx) error {
	request := new(UpdatePaymentStatusReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(web.UnprocessableEntity(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.BadRequest(err))
	}

	req := <-s.registrantRepo.Query.GetByID(request.RegisterID)
	if req.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	registrantData := req.Result.(ReadResult)
	if registrantData.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFound(err))
	}

	var status PaymentStatus
	if request.PaymentType == "credit_card" && request.PaymentStatus == "capture" && request.FraudStatus == "accept" {
		status = PaymentStatusPaid
	} else if request.PaymentStatus == "settlement" {
		status = PaymentStatusPaid
	} else if request.PaymentStatus == "deny" || request.PaymentStatus == "expire" || request.PaymentStatus == "cancel" {
		status = PaymentStatusCancel
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	err = <-s.registrantRepo.Cmd.UpdateStatus(request.RegisterID, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(web.Ok(nil))
}

func (s *Server) HandleRegistrantLogin(c *fiber.Ctx) error {
	request := new(LoginReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(web.UnprocessableEntity(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.BadRequest(err))
	}

	req := <-s.registrantRepo.Query.GetByUsernameAndPassword(request.Username, request.Password)
	if req.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	registrantData := req.Result.(ReadResult)
	if registrantData.Password == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFound(errors.New("registrant not found")))
	}

	accessToken, err := s.tokenMaker.CreateToken(
		registrantData.Username,
		config.AccessTokenDuration,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.NotFound(err))
	}

	rsp := LoginResponse{
		AccessToken: accessToken,
		Registrant:  registrantData,
	}

	return c.Status(fiber.StatusOK).JSON(web.Ok(rsp))
}
