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
	product, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return product, nil
}
