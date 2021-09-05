package server

import (
	"github.com/SemmiDev/go-pmb/internal/common/config"
	"github.com/SemmiDev/go-pmb/internal/payment"
	"github.com/SemmiDev/go-pmb/internal/registrant/command"
	registrantCmdMySql "github.com/SemmiDev/go-pmb/internal/registrant/command/mysql"
	"github.com/SemmiDev/go-pmb/internal/registrant/query"
	registrantQueryMySql "github.com/SemmiDev/go-pmb/internal/registrant/query/mysql"
	"github.com/gofiber/fiber/v2"
)

type RegistrantServer struct {
	RegistrantRepo  command.RegistrantCommand
	RegistrantQuery query.RegistrantQuery
	Midtrans        payment.IMidtrans
}

func NewRegistrantServer() *RegistrantServer {
	registrantServer := &RegistrantServer{}

	mySqlConnection := config.MySqlDB

	registrantServer.Midtrans = payment.NewMidtrans()
	registrantServer.RegistrantQuery = registrantQueryMySql.NewRegistrantMySqlQuery(mySqlConnection)
	registrantServer.RegistrantRepo = registrantCmdMySql.NewRegistrantCommandMysql(mySqlConnection)

	return registrantServer
}

func (r *RegistrantServer) Mount(app *fiber.App) {
	app.Post("/api/registrant/register", r.HandleRegisterRegistrant)
	app.Put("/api/registrant/payment_status", r.HandleUpdateRegisterPaymentStatus)
}
