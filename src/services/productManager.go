package services

import "store/src/models"

// ProductManager provides CRUD access to product entities
type ProductManager struct {
}

// NewProductManager creates a new instance of a product manager
func NewProductManager() *ProductManager {
	return &ProductManager{}
}

// GetList returns a slice of User entities
func (productManager *ProductManager) GetList() ([]models.Product, error) {
	product1 := models.NewProduct("Product 1", "Description of product 1", 19.99)
	return []models.Product{product1}, nil
}

// Get returns a product that matches by id
func (productManager *ProductManager) Get(id int) (models.Product, error) {
	product := models.NewProduct("name", "description", 19.99)
	return product, nil
}

// Save upserts a product in the data store
func (productManager *ProductManager) Save(product models.Product) (models.Product, error) {
	updatedProduct := models.NewProduct(product.Name, product.Description, product.Price)

	return updatedProduct, nil
}

// Delete removes an entity by id from the data store
func (productManager *ProductManager) Delete(id int) (bool, error) {
	return true, nil
}
