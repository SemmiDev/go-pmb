package registrant

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/common/config"
	"github.com/SemmiDev/go-pmb/common/token"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	registrantRepo Repository
	midtrans       IMidtrans
	tokenMaker     token.Maker
	router         *fiber.App
}

func NewServer() (*Server, error) {
	mySqlConnection := config.MySQLConnect()

	midtrans := NewMidtrans()
	command := NewCommandMysql(mySqlConnection)
	query := NewMySqlQuery(mySqlConnection)

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		registrantRepo: Repository{
			Cmd:   command,
			Query: query,
		},
		midtrans:   midtrans,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

func (s *Server) Mount(router *fiber.App) {
	router.Post("/api/registrant/register", s.HandleRegisterRegistrant)
	router.Post("/api/registrant/login", s.HandleRegistrantLogin)
	router.Put("/api/registrant/payment_status", s.HandleUpdateRegisterPaymentStatus)

	s.router = router
}
