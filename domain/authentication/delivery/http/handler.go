package http

import (
	"go-boilerplate/domain/authentication"
	"go-boilerplate/models"
	"go-boilerplate/shared/utils"

	"github.com/labstack/echo"
)

type authenticationHandler struct {
	authenticationService authentication.Service
}

func NewAuthenticationHandler(e *echo.Echo, service authentication.Service) {
	handler := authenticationHandler{
		authenticationService: service,
	}

	e.POST("/create-session", handler.CreateSession)
}

func (h authenticationHandler) CreateSession(c echo.Context) error {
	var createSession models.CreateSessionRequest
	if err := c.Bind(&createSession); err != nil {
		return utils.ResponseJSON(c, 400, nil, err.Error())
	}

	if err := c.Validate(&createSession); err != nil {
		return utils.ResponseJSON(c, 400, nil, err.Error())
	}

	result, err := h.authenticationService.GenerateToken(createSession)
	if err != nil {
		return utils.ResponseJSON(c, 400, nil, err.Error())
	}

	return utils.ResponseJSON(c, 200, result, "")
}
