package context

import (
	"go-boilerplate/models"
	"go-boilerplate/shared/config"
	"go-boilerplate/shared/validator"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type ApplicationContext struct {
	echo.Context
}

func (c *ApplicationContext) GetConfig(key string) interface{} {
	config.LoadConfig()
	val := viper.Get(key)
	return val
}

func (c *ApplicationContext) Validate(i interface{}) error {
	val := validator.DefaultValidator()
	return val.Validate(i)
}

func (c *ApplicationContext) ResponseJSON(status int, data interface{}, message string) error {
	return c.JSON(status, &models.Response{
		Data:    data,
		Message: message,
		Code:    status,
	})
}
