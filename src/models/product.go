package models

var empty Product = Product{}

// Product stores product information
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

// NewProduct creates a new instance of a product
func NewProduct(name, description string, price float32) *Product {
	return &Product{
		// TODO automatically incremend id
		ID:          1,
		Name:        name,
		Description: description,
		Price:       price,
	}
}

// Empty returns an empty product
func Empty() *Product {
	return &empty
}
