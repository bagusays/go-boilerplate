package http

import (
	"go-boilerplate/domain/authentication"
	"go-boilerplate/models"
	"go-boilerplate/shared/context"

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
	ct := c.(*context.CustomApplicationContext)

	var createSession models.CreateSessionRequest
	if err := ct.Bind(&createSession); err != nil {
		return ct.ResponseJSON(400, nil, err.Error())
	}

	if err := ct.Validate(&createSession); err != nil {
		return ct.ResponseJSON(400, nil, err.Error())
	}

	result, err := h.authenticationService.GenerateToken(createSession)
	if err != nil {
		return ct.ResponseJSON(400, nil, err.Error())
	}

	return ct.ResponseJSON(200, result, "")
}
