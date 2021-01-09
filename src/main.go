package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wbreza/go-store/api/controllers"
)

func main() {
	router := mux.NewRouter()
	controllers.RegisterControllers(router)
	fmt.Println("Running server on port 3000")
	http.ListenAndServe(":3000", router)
}
