package main

import (
	healthCheck "go-boilerplate/domain/health-check/delivery/http"

	productHandler "go-boilerplate/domain/products/delivery/http"
	productRepository "go-boilerplate/domain/products/repository"
	productService "go-boilerplate/domain/products/service"

	authenticationHandler "go-boilerplate/domain/authentication/delivery/http"
	authenticationRepository "go-boilerplate/domain/authentication/repository"
	authenticationService "go-boilerplate/domain/authentication/service"

	"go-boilerplate/shared/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	config.LoadConfig()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","method":"${method}","uri":"${uri}","status":${status},` +
			`"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	}))

	healthCheck.NewHealthCheckHandler(e)

	productRepo := productRepository.NewProductsRepository()
	productService := productService.NewProductService(productRepo)
	productHandler.NewProductHandler(e, productService)

	authenticationRepo := authenticationRepository.NewAuthenticationRepository()
	authenticationService := authenticationService.NewAuthenticationService(authenticationRepo)
	authenticationHandler.NewAuthenticationHandler(e, authenticationService)

	e.Logger.Fatal(e.Start(":1323"))
}
