package service

import (
	"go-boilerplate/domain/products"
	"go-boilerplate/models"
)

type serviceHandler struct {
	productRepo products.Repository
}

func NewProductService(productRepo products.Repository) products.Service {
	return &serviceHandler{
		productRepo: productRepo,
	}
}

func (s serviceHandler) GetProducts() (*models.Products, error) {
	return s.productRepo.GetProducts()
}
