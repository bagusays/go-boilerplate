package http

import (
	"go-boilerplate/domain/products"
	"go-boilerplate/shared/middleware"
	"go-boilerplate/shared/utils"

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
	result := h.productsService.GetProducts()
	return utils.ResponseJSON(c, 200, result, "")
}
