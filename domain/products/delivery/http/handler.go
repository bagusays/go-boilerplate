package http

import (
	"go-boilerplate/domain/products"
	"go-boilerplate/shared/utils"

	"github.com/labstack/echo"
)

type productsHandler struct {
	productsService products.Service
}

func NewProductHandler(e *echo.Echo, service products.Service) productsHandler {
	handler := productsHandler{
		productsService: service,
	}

	e.GET("/", handler.GetProducts)

	return handler
}

func (h productsHandler) GetProducts(c echo.Context) error {
	result, err := h.productsService.GetProducts()
	if err != nil {
		return utils.ResponseJSON(c, 500, nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, result, "")
}
