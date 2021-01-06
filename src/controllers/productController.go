package controllers

import (
	"encoding/json"
	"net/http"
	"store/src/services"
)

// ProductController exposes actions on the products API
type ProductController struct {
	productManager services.ProductManager
}

// NewProductController creates a new instance of a product controller
func NewProductController() *ProductController {
	controller := ProductController{
		productManager: *services.NewProductManager(),
	}

	return &controller
}

func (controller *ProductController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	controller.GetList(writer, request)
}

// GetList gets a list of products
func (controller *ProductController) GetList(writer http.ResponseWriter, request *http.Request) {
	products, err := controller.productManager.GetList()

	if err != nil {
		writer.WriteHeader(500)
	}

	encoder := json.NewEncoder(writer)
	encoder.Encode(products)
}
