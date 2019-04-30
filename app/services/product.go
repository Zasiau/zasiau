package services

import (
	"database/sql"
	"net/http"

	"github.com/dongri/gonion/app/middlewares/postgres"
	"github.com/dongri/gonion/app/models"
)

// GetProducts ...
func GetProducts(r *http.Request) ([]models.Product, error) {
	dbmap := postgres.GetDbMap(r)
	products, err := models.ProductList(dbmap)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return products, nil
}

// GetProductByID ...
func GetProductByID(r *http.Request, ID uint64) (*models.Product, error) {
	dbmap := postgres.GetDbMap(r)
	product, err := models.ProductFindByID(dbmap, ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// CreateProduct ...
func CreateProduct(r *http.Request, name string, price uint32) error {
	dbmap := postgres.GetDbMap(r)
	product := new(models.Product)
	product.Name = name
	product.Price = price
	return product.Insert(dbmap)
}

// UpdateProductByID ...
func UpdateProductByID(r *http.Request, ID uint64, name string, price uint32) error {
	dbmap := postgres.GetDbMap(r)
	product, err := models.ProductFindByID(dbmap, ID)
	if err != nil {
		return err
	}
	product.Name = name
	product.Price = price
	return product.Update(dbmap)
}
