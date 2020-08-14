package repository

import (
	"go-boilerplate/domain/products"
)

type repoHandler struct {
}

// NewVoucherRepository ....
func NewProductsRepository() products.Repository {
	return &repoHandler{}
}

func (r *repoHandler) GetProducts() interface{} {
	return map[string]interface{}{
		"name": "Chocolatos",
	}
}
