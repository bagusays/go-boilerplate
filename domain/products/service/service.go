package service

import (
	"go-boilerplate/domain/products"
)

type serviceHandler struct {
	productRepo products.Repository
}

func NewProductService(productRepo products.Repository) products.Service {
	return &serviceHandler{
		productRepo: productRepo,
	}
}

func (s serviceHandler) GetProducts() interface{} {
	return s.productRepo.GetProducts()
}
