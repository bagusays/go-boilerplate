package http

import (
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckHandler struct{}

func NewHealthCheckHandler(e *echo.Echo) *healthCheckHandler {
	handler := healthCheckHandler{}

	e.GET("/health-check", handler.HealthCheckHandler)

	return &handler
}

func (h healthCheckHandler) HealthCheckHandler(c echo.Context) error {
	return utils.ResponseJSON(c, http.StatusOK, constants.StatusSuccess, "ok", "")
}
