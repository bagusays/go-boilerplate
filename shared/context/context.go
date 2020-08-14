package context

import (
	"fmt"
	"go-boilerplate/models"
	"go-boilerplate/shared/config"
	"go-boilerplate/shared/validator"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type CustomApplicationContext struct {
	echo.Context
}

func (c *CustomApplicationContext) GetConfig(key string) interface{} {
	config.LoadConfig()
	val := viper.Get(key)
	fmt.Println(val)
	return val
}

func (c *CustomApplicationContext) Validate(i interface{}) error {
	val := validator.DefaultValidator()
	return val.Validate(i)
}

func (c *CustomApplicationContext) ResponseJSON(status int, data interface{}, message string) error {
	return c.JSON(status, &models.Response{
		Data:    data,
		Message: message,
		Code:    status,
	})
}
