package controllers

import (
	"github.com/gorilla/mux"
)

// RegisterControllers registers all controllers
func RegisterControllers(router *mux.Router) {
	NewProductController(router)
}
