package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/app"
	"github.com/SemmiDev/fiber-go-clean-arch/api/app/configuration"
)

func main() {
	appConfig := configuration.App()
	app.Run(appConfig)
	// app.RunServerWithGracefulShutdown()
}
