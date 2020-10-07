package main

import (
	"fmt"
	"go-boilerplate/cmd"
	healthCheck "go-boilerplate/domain/health-check/delivery/http"

	productHandler "go-boilerplate/domain/products/delivery/http"
	productRepository "go-boilerplate/domain/products/repository"
	productService "go-boilerplate/domain/products/service"

	authenticationHandler "go-boilerplate/domain/authentication/delivery/http"
	authenticationRepository "go-boilerplate/domain/authentication/repository"
	authenticationService "go-boilerplate/domain/authentication/service"

	"go-boilerplate/shared/config"
	"go-boilerplate/shared/database"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()
	cmd.LoadCMDConfiguration()

	e := echo.New()

	mysqlConn, err := database.OpenMysqlConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open mysql connection: %s", err.Error())
		log.Errorf(msgError)
		panic(msgError)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","method":"${method}","uri":"${uri}","status":${status},` +
			`"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	}))

	healthCheck.NewHealthCheckHandler(e)

	productRepo := productRepository.NewProductsRepository(mysqlConn)
	productService := productService.NewProductService(productRepo)
	productHandler.NewProductHandler(e, productService)

	authenticationRepo := authenticationRepository.NewAuthenticationRepository()
	authenticationService := authenticationService.NewAuthenticationService(authenticationRepo)
	authenticationHandler.NewAuthenticationHandler(e, authenticationService)

	e.Logger.Fatal(e.Start(":1323"))
}
