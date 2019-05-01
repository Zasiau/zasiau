package models

import (
	"fmt"
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Product ...
type Product struct {
	Base
	Name  string `db:"name" json:"name"`
	Price uint32 `db:"price" json:"price"`
}

// const ...
const (
	ProductColumns = "id, name, price, created, updated"
)

// Insert ...
func (s *Product) Insert(exec gorp.SqlExecutor) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	return exec.Insert(s)
}

// Update ...
func (s *Product) Update(exec gorp.SqlExecutor) error {
	s.Updated = time.Now().UTC()
	_, err := exec.Update(s)
	return err
}

// Delete ...
func (s *Product) Delete(exec gorp.SqlExecutor) error {
	s.Updated = time.Now().UTC()
	_, err := exec.Delete(s)
	return err
}

// ProductFindByID ...
func ProductFindByID(exec gorp.SqlExecutor, ID uint64) (*Product, error) {
	product := new(Product)
	query := fmt.Sprintf("SELECT %s FROM products WHERE id = $1", ProductColumns)
	if err := exec.SelectOne(&product, query, ID); err != nil {
		return nil, err
	}
	return product, nil
}

// ProductList ...
func ProductList(exec gorp.SqlExecutor) ([]Product, error) {
	products := []Product{}
	query := fmt.Sprintf("SELECT %s FROM products", ProductColumns)
	if _, err := exec.Select(&products, query); err != nil {
		return nil, err
	}
	return products, nil
}

// DeleteProductFindByID ...
func DeleteProductFindByID(exec gorp.SqlExecutor, ID uint64) (*Product, error) {
	product := new(Product)
	query := fmt.Sprintf("DELETE FROM products WHERE id = $1")
	if err := exec.SelectOne(&product, query, ID); err != nil {
		return nil, err
	}
	return product, nil
}
