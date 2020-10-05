package http

import (
	"fmt"
	"go-boilerplate/shared/config"
	"go-boilerplate/shared/utils"

	"github.com/labstack/echo"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler(e *echo.Echo) healthCheckHandler {
	handler := healthCheckHandler{}

	e.GET("/health-check", handler.HealthCheckHandler)

	return handler
}

func (h healthCheckHandler) HealthCheckHandler(c echo.Context) error {
	fmt.Println(config.GetConfig().GetJWTSecretKey())
	fmt.Println(config.GetConfig().GetServiceName())

	return utils.ResponseJSON(c, 200, "ok", "")
}
