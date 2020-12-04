package http

import (
	"go-boilerplate/domain/products"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/utils"
	"net/http"

	"github.com/labstack/echo/v4"
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
		return utils.ResponseJSON(c, http.StatusInternalServerError, errors.MapErrToStatusCode(err), nil, err.Error())
	}
	return utils.ResponseJSON(c, constants.StatusSuccess, 200, result, "")
}
