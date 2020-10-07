package repository

import (
	"go-boilerplate/domain/products"
	"go-boilerplate/models"

	"gorm.io/gorm"
)

type repoHandler struct {
	mysqlConn *gorm.DB
}

// NewVoucherRepository ....
func NewProductsRepository(mysqlConn *gorm.DB) products.Repository {
	return &repoHandler{
		mysqlConn: mysqlConn,
	}
}

func (r repoHandler) GetProducts() (*models.Products, error) {
	var product models.Products
	res := r.mysqlConn.
		Table("products").
		First(&product)
	return &product, res.Error
}
