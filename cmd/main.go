package main

import (
	"github.com/silapdev/go_i18n_test/internal/service"
	"github.com/silapdev/go_i18n_test/internal/transport/rest"
	"github.com/silapdev/go_i18n_test/pkg/locales"
	"log"
)

func main() {

	services := service.NewService()

	bundle := locales.NewBundle()

	handler := rest.NewHandler(services, bundle)

	app := handler.InitRoutes()

	port := ":3001"

	err := app.Listen(port)

	if err != nil {
		log.Fatal()
	}

}
