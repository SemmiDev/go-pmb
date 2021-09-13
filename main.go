package main

import (
	"database/sql"
	"fmt"
	"github.com/SemmiDev/go-pmb/pkg/common/config"
	"github.com/SemmiDev/go-pmb/pkg/payment"
	"github.com/SemmiDev/go-pmb/pkg/registrant"
	storedb "github.com/SemmiDev/go-pmb/pkg/stores/sql"
	"github.com/gofiber/fiber/v2"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load configuration file.
	config.Load()

	// creates a new fiber instance.
	app := fiber.New()

	// mysql data source name.
	DSN := config.MysqlUser + ":" + config.MysqlPassword + "@(" + config.MysqlHost + ":" + config.MysqlPort + ")/" + config.MysqlDbname + "?parseTime=true&clientFoundRows=true"

	// open mysql db.
	mySqlDb, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setup mysql.
	sqlDb := storedb.SetupMySql(mySqlDb)

	// setup registrant mysql repository.
	registrantRepo := registrant.NewMySqlRepository(sqlDb)

	// setup midtrans payment gateway.
	mid := payment.NewMidtrans()

	// setup registrant service.
	registrantService := registrant.NewService(registrantRepo, mid)

	// setup registrant server.
	registrantServer := registrant.NewServer(registrantService)

	// add check health endpoint.
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("okay")
	})

	router := app.Group("api")
	// register registrant endpoints.
	registrantServer.Mount(router)

	// run app.
	fmt.Printf("running on http://localhost%v", config.AppPort)
	log.Fatal(app.Listen(config.AppPort))
}
