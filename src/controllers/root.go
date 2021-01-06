package controllers

import "net/http"

// RegisterControllers registers all controllers
func RegisterControllers() {
	productController := NewProductController()

	http.Handle("/products", productController)
}
