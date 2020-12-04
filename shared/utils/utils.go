package utils

import (
	"go-boilerplate/models"

	"github.com/labstack/echo/v4"
)

func ResponseJSON(c echo.Context, statusCode int, httpStatus int, data interface{}, message string) error {
	return c.JSON(httpStatus, models.Response{
		Data:    data,
		Message: message,
		Code:    statusCode,
	})
}
