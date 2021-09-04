package server

import (
	"github.com/SemmiDev/go-pmb/internal/database"
	"github.com/SemmiDev/go-pmb/pkg/helper"
	"github.com/SemmiDev/go-pmb/pkg/payment"
	"github.com/SemmiDev/go-pmb/pkg/registrant/domain"
	"github.com/SemmiDev/go-pmb/pkg/registrant/errors"
	"github.com/SemmiDev/go-pmb/pkg/registrant/query"
	registrantQueryMySql "github.com/SemmiDev/go-pmb/pkg/registrant/query/mysql"
	"github.com/SemmiDev/go-pmb/pkg/registrant/repository"
	registrantRepoMySql "github.com/SemmiDev/go-pmb/pkg/registrant/repository/mysql"
	"github.com/SemmiDev/go-pmb/pkg/registrant/storage"
	"github.com/SemmiDev/go-pmb/pkg/web"
	"github.com/gofiber/fiber/v2"
	"github.com/myesui/uuid"
)

type RegistrantServer struct {
	RegistrantRepo  repository.RegistrantRepository
	RegistrantQuery query.RegistrantQuery
	Midtrans        payment.IMidtrans
}

func NewRegistrantServer() *RegistrantServer {
	registrantServer := &RegistrantServer{}

	mySql := database.MySQL{}
	mySqlConnection, err := mySql.Connect()
	if err != nil {
		panic(err)
	}

	registrantServer.Midtrans = payment.NewMidtrans()
	registrantServer.RegistrantQuery = registrantQueryMySql.NewRegistrantMySqlQuery(mySqlConnection)
	registrantServer.RegistrantRepo = registrantRepoMySql.NewRegistrantRepositoryMysql(mySqlConnection)

	return registrantServer
}

func (r *RegistrantServer) Mount(app *fiber.App) {
	app.Post("/api/registrant/register", r.Register)
	app.Put("/api/registrant/payment_status", r.UpdateRegisterPaymentStatus)
}

func (r *RegistrantServer) Register(c *fiber.Ctx) error {
	request := new(domain.CreateRegistrantReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(web.UnprocessableEntity(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.BadRequest(err))
	}

	result, err := r.registerRegistrant(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	return c.Status(fiber.StatusCreated).JSON(web.Created(result))
}

func (r *RegistrantServer) registerRegistrant(req *domain.CreateRegistrantReq) (res *domain.RegisterResponse, err error) {
	registrant := domain.Registrant{}

	password := helper.GeneratePassword()
	registrant.New(req.Name, req.Email, req.Phone, password)

	if req.Program == domain.ProgramS1D3D4 {
		registrant.Program = domain.ProgramS1D3D4
		registrant.Bill = domain.BillS1D3D4
	} else if req.Program == domain.ProgramS2 {
		registrant.Program = domain.ProgramS2
		registrant.Bill = domain.BillS2
	} else {
		return nil, errors.RegistrantError{Code: errors.RegistrantErrorProgramNotSupportedCode}
	}

	paymentPayload := payment.Payload{
		ID:     uuid.NewV4().String(),
		Amount: registrant.Bill.Value(),
	}

	paymentURL, err := r.Midtrans.GetPaymentURL(&paymentPayload, &registrant)
	if err != nil {
		return nil, err
	}
	registrant.PaymentURL = paymentURL

	err = <-r.RegistrantRepo.Save(&registrant)
	if err != nil {
		return nil, err
	}

	return registrant.ToRegistrantResponse(password, paymentPayload.AmountFormatIDR()), nil
}

func (r *RegistrantServer) UpdateRegisterPaymentStatus(c *fiber.Ctx) error {
	request := new(domain.UpdatePaymentStatusReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(web.UnprocessableEntity(err))
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.BadRequest(err))
	}

	registrant := <-r.RegistrantQuery.GetByID(request.RegisterID)
	if registrant.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	registrantData := registrant.Result.(storage.RegistrantResult)
	if registrantData.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.NotFound(err))
	}

	var status domain.PaymentStatus
	if request.PaymentType == "credit_card" && request.PaymentStatus == "capture" && request.FraudStatus == "accept" {
		status = domain.PaymentStatusPaid
	} else if request.PaymentStatus == "settlement" {
		status = domain.PaymentStatusPaid
	} else if request.PaymentStatus == "deny" || request.PaymentStatus == "expire" || request.PaymentStatus == "cancel" {
		status = domain.PaymentStatusCancel
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	err = <-r.RegistrantRepo.UpdateStatus(request.RegisterID, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.InternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(web.Ok(nil))
}
