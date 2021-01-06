package services

import "store/src/models"

var cache = make(map[int]*models.Product)

// ProductManager provides CRUD access to product entities
type ProductManager struct {
}

// NewProductManager creates a new instance of a product manager
func NewProductManager() *ProductManager {
	return &ProductManager{}
}

// GetList returns a slice of User entities
func (productManager *ProductManager) GetList() ([]*models.Product, error) {
	products := []*models.Product{}
	for _, value := range cache {
		products = append(products, value)
	}

	return products, nil
}

// Get returns a product that matches by id
func (productManager *ProductManager) Get(id int) (*models.Product, error) {
	return cache[id], nil
}

// Save inserts or updates a product in the data store
func (productManager *ProductManager) Save(product *models.Product) (*models.Product, error) {
	if product.ID == 0 {
		product.ID = GetNewID()
	}

	cache[product.ID] = product

	return product, nil
}

// Delete removes an entity by id from the data store
func (productManager *ProductManager) Delete(id int) (bool, error) {
	product := cache[id]
	if product == nil {
		return false, nil
	}

	delete(cache, id)
	return true, nil
}

// GetNewID returns the next id for storing stuff
func GetNewID() int {
	maxKey := 0
	for key := range cache {
		if key > maxKey {
			maxKey = key
		}
	}

	return maxKey + 1
}
