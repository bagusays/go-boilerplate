package products

import "go-boilerplate/models"

type Service interface {
	GetProducts() (*models.Products, error)
}

type Repository interface {
	GetProducts() (*models.Products, error)
}
