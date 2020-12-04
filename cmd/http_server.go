package cmd

import (
	"context"
	"fmt"
	healthCheck "go-boilerplate/domain/health-check/delivery/http"
	productHandler "go-boilerplate/domain/products/delivery/http"
	productRepository "go-boilerplate/domain/products/repository"
	productService "go-boilerplate/domain/products/service"
	"go-boilerplate/shared/config"
	"go-boilerplate/shared/infrastructures/database"
	"go-boilerplate/shared/log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type ServerDependencies struct {
	IsReady    chan bool
	IsShutdown chan bool
	Config     *config.Config
}

func NewServer(server ServerDependencies) *ServerDependencies {
	return &server
}

func (i *ServerDependencies) StartServer() {
	log := log.NewLogger(i.Config.Log)

	mysqlConn, err := database.OpenMysqlConn(i.Config.Database)
	if err != nil {
		panic("Failed to open mysql connection: " + err.Error())
	}

	server := echo.New()

	healthCheck.NewHealthCheckHandler(server)

	productRepo := productRepository.NewProductsRepository(mysqlConn)
	productSvc := productService.NewProductService(productRepo)
	productHandler.NewProductHandler(server, productSvc)

	errCh := make(chan error)
	go func() {
		if err := server.Start(fmt.Sprintf(":%s", i.Config.Apps.Port)); err != nil {
			errCh <- err
			log.Error("Failed to shutdown server", err)
		}
	}()

	i.IsReady <- true

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
		close(i.IsShutdown)
	}()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown server", err)
	}
}
