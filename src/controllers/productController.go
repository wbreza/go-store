package controllers

import (
	"net/http"
	"store/src/models"
	"store/src/services"

	"github.com/gorilla/mux"
)

// ProductController exposes actions on the products API
type ProductController struct {
	Controller
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

	controller.WriteJSON(writer, products)
}

// GetByID gets a product with the specified id
func (controller *ProductController) GetByID(writer http.ResponseWriter, request *http.Request) {
	productID, err := controller.GetRequestParamAsInt(request, "productId")

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

	controller.WriteJSON(writer, product)
}

// Create saves a new product to the data store
func (controller *ProductController) Create(writer http.ResponseWriter, request *http.Request) {
	product := models.Product{}
	err := controller.ParseJSONBody(request, &product)

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
	controller.WriteJSON(writer, updatedProduct)
}

// Update updates a current product with the matching id
func (controller *ProductController) Update(writer http.ResponseWriter, request *http.Request) {
	productID, err := controller.GetRequestParamAsInt(request, "productId")

	if err != nil {
		writer.WriteHeader(400)
		return
	}

	product := models.Product{}
	err = controller.ParseJSONBody(request, &product)

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
	productID, err := controller.GetRequestParamAsInt(request, "productId")

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
	} else {
		writer.WriteHeader(404)
	}
}
