package http

import (
	"go-boilerplate/domain/products"
	"go-boilerplate/shared/context"
	"go-boilerplate/shared/middleware"

	"github.com/labstack/echo"
)

type productsHandler struct {
	productsService products.Service
}

func NewProductHandler(e *echo.Echo, service products.Service) {
	handler := productsHandler{
		productsService: service,
	}

	e.GET("/", middleware.Authentication(handler.GetProducts))
}

func (h productsHandler) GetProducts(c echo.Context) error {
	ct := c.(*context.CustomApplicationContext)
	result := h.productsService.GetProducts()
	return ct.ResponseJSON(200, result, "")
}
