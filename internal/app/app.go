package app

import (
	"os"
	"os/signal"
	"sideq/config"
	"sideq/internal/controller/router"
	"sideq/internal/usecase"
	"sideq/internal/usecase/repo"
	"sideq/pkg/db"
	"sideq/pkg/httpserver"
	"sideq/pkg/logger"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) error {
	l := logger.New(cfg.Log.Level)
	database, err := db.New(cfg.GORM.URL)
	if err != nil {
		return err
	}

	exampleUseCase := usecase.New(repo.New(database))

	handler := gin.New()
	router.NewRouter(handler, l, exampleUseCase)
	httpServer := httpserver.New(handler, cfg.HTTP.Address, cfg.HTTP.Timeout)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	err = httpServer.Shutdown()
	if err != nil {
		return err
	}

	return nil
}
