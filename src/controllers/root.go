package controllers

import (
	"github.com/gorilla/mux"
)

// RegisterControllers registers all controllers
func RegisterControllers(router *mux.Router) {
	NewProductController(router)
	// TODO how come it's ok not to set this to something?
}
