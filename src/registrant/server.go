package registrant

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/src/common/config"
	"github.com/SemmiDev/go-pmb/src/common/token"
	"github.com/SemmiDev/go-pmb/src/registrant/interfaces"
	"github.com/SemmiDev/go-pmb/src/registrant/payments"
	"github.com/SemmiDev/go-pmb/src/registrant/repositories"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	registrantRepo interfaces.IRepository
	midtrans       payments.IMidtrans
	tokenMaker     token.Maker
	router         *fiber.App
}

func NewServer() (*Server, error) {
	mySqlConnection := config.MySQLConnect()

	midtrans := payments.NewMidtrans()
	repo := repositories.NewRegistrantRepository(mySqlConnection)

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		registrantRepo: repo,
		midtrans:       midtrans,
		tokenMaker:     tokenMaker,
	}

	return server, nil
}

func (s *Server) Mount(router *fiber.App) {
	router.Post("/api/registrant/register", s.HandleRegisterRegistrant)
	router.Post("/api/registrant/login", s.HandleRegistrantLogin)
	router.Put("/api/registrant/payment_status", s.HandleUpdateRegisterPaymentStatus)

	s.router = router
}
