package main

import (
	"fmt"
	"log"
	"net/http"

	"contabi-be/config"
	"contabi-be/controller"
	"contabi-be/middleware"
	"contabi-be/router"
	"contabi-be/service/database"
	"contabi-be/usecase"

	"github.com/sirupsen/logrus"
)

func main() {
	// load env vars
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	//creates logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(logrus.InfoLevel)

	// creates service instances
	dbs, err := database.NewDatabaseService(cfg)
	if err != nil {
		log.Fatalf("Error al crear el servicio de base de datos: %v", err)
	}

	// Creates instances of the services
	ls := database.NewLoginService(dbs.DB)
	us := database.NewUsersService(dbs.DB)
	cs := database.NewClientsService(dbs.DB)
	ms := database.NewMenusService(dbs.DB)

	// creates instances of usecase
	lu := usecase.NewLoginUseCase(ls)
	uu := usecase.NewUsersUseCase(us)
	cu := usecase.NewClientsUseCase(cs)
	mu := usecase.NewMenusUseCase(ms)

	// creates instances of controller
	lc := controller.NewLoginController(lu, logger)
	uc := controller.NewUsersController(uu, logger)
	cc := controller.NewClientsController(cu, logger)
	mc := controller.NewMenusController(mu, logger)
	mw := middleware.New(lu)

	// creates router instance
	rr := router.NewRouter(
		lc,
		uc,
		cc,
		mc,
		mw,
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: rr,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Fatal error starting the server: %s", err)
	}
}
