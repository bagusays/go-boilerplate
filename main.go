package main

import (
	productHandler "go-boilerplate/domain/products/delivery/http"
	productRepository "go-boilerplate/domain/products/repository"
	productService "go-boilerplate/domain/products/service"

	authenticationHandler "go-boilerplate/domain/authentication/delivery/http"
	authenticationRepository "go-boilerplate/domain/authentication/repository"
	authenticationService "go-boilerplate/domain/authentication/service"

	"go-boilerplate/shared/context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","method":"${method}","uri":"${uri}","status":${status},` +
			`"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.ApplicationContext{c}
			return next(cc)
		}
	})

	productRepo := productRepository.NewProductsRepository()
	productService := productService.NewProductService(productRepo)
	productHandler.NewProductHandler(e, productService)

	authenticationRepo := authenticationRepository.NewAuthenticationRepository()
	authenticationService := authenticationService.NewAuthenticationService(authenticationRepo)
	authenticationHandler.NewAuthenticationHandler(e, authenticationService)

	e.Logger.Fatal(e.Start(":1323"))
}
