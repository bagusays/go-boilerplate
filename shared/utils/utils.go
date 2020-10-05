package utils

import (
	"go-boilerplate/models"

	"github.com/labstack/echo"
)

func ResponseJSON(c echo.Context, status int, data interface{}, message string) error {
	return c.JSON(status, &models.Response{
		Data:    data,
		Message: message,
		Code:    status,
	})
}
