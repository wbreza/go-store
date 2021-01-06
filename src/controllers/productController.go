package controllers

import (
	"encoding/json"
	"net/http"
	"store/src/models"
	"store/src/services"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductController exposes actions on the products API
type ProductController struct {
	productManager services.ProductManager
	router         mux.Router
}

// NewProductController creates a new instance of a product controller
func NewProductController(router *mux.Router) *ProductController {
	controller := ProductController{
		productManager: *services.NewProductManager(),
		router:         *router,
	}

	router.HandleFunc("/products", controller.GetList).Methods(http.MethodGet)
	router.HandleFunc("/products", controller.Create).Methods(http.MethodPost)
	router.HandleFunc("/products/{productId}", controller.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/products/{productId}", controller.Update).Methods(http.MethodPut)
	router.HandleFunc("/products/{productId}", controller.DeleteByID).Methods(http.MethodDelete)

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

// GetByID gets a product with the specified id
func (controller *ProductController) GetByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	productID, err := strconv.Atoi(vars["productId"])

	if err != nil {
		writer.WriteHeader(400)
		return
	}

	product, err := controller.productManager.Get(productID)

	if err != nil {
		writer.WriteHeader(500)
		return
	}

	if product == nil {
		writer.WriteHeader(404)
		return
	}

	encoder := json.NewEncoder(writer)
	encoder.Encode(product)
}

// Create saves a new product to the data store
func (controller *ProductController) Create(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	product := models.Product{}
	err := decoder.Decode(&product)
	if err != nil {
		writer.WriteHeader(400)
		return
	}

	updatedProduct, err := controller.productManager.Save(&product)
	if err != nil {
		writer.WriteHeader(500)
		return
	}

	writer.WriteHeader(201)
	encoder := json.NewEncoder(writer)
	encoder.Encode(updatedProduct)
}

// Update updates a current product with the matching id
func (controller *ProductController) Update(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	productID, err := strconv.Atoi(vars["productId"])

	decoder := json.NewDecoder(request.Body)
	product := models.Product{}
	err = decoder.Decode(&product)
	if err != nil {
		writer.WriteHeader(400)
		return
	}

	product.ID = productID

	if err != nil {
		writer.WriteHeader(400)
		return
	}

	controller.productManager.Save(&product)
}

// DeleteByID removes the product with the specified id
func (controller *ProductController) DeleteByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	productID, err := strconv.Atoi(vars["productId"])

	if err != nil {
		writer.WriteHeader(400)
		return
	}

	isDeleted, err := controller.productManager.Delete(productID)
	if err != nil {
		writer.WriteHeader(500)
		return
	}

	if isDeleted {
		writer.WriteHeader(204)
		return
	} else {
		writer.WriteHeader(404)
		return
	}
}
