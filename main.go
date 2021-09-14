package main

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/pkg/common/config"
	"github.com/SemmiDev/go-pmb/pkg/common/database"
	"github.com/SemmiDev/go-pmb/pkg/payment"
	"github.com/SemmiDev/go-pmb/pkg/registrant/controller"
	"github.com/SemmiDev/go-pmb/pkg/registrant/repository"
	"github.com/SemmiDev/go-pmb/pkg/registrant/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// load environments variable.
	config.LoadEnv()

	// creates a new fiber instance.
	app := fiber.New()

	// mysql data source name.
	DSN := config.MysqlUser + ":" + config.MysqlPassword + "@(" + config.MysqlHost + ":" + config.MysqlPort + ")/" + config.MysqlDbname + "?parseTime=true&clientFoundRows=true"

	// setup mysql.
	mysqlDb := database.NewSqlDb(DSN).Open()

	// setup registrant mysql repository.
	registrantRepo := repository.NewMySqlRepository(mysqlDb)

	// setup midtrans payment gateway.
	mid := payment.NewMidtrans()

	// setup registrant service.
	registrantService := service.NewService(registrantRepo, mid)

	// setup registrant controller.
	registrantServer := controller.NewController(registrantService)

	// add check health endpoint.
	router := app.Group("api")
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("okay")
	})

	// register registrant endpoints.
	registrantServer.Mount(router)

	// run app.
	fmt.Printf("running on http://localhost%v", config.AppPort)
	log.Fatal(app.Listen(config.AppPort))
}
