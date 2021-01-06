package main

import (
	"fmt"
	"net/http"
	"store/src/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	controllers.RegisterControllers(router)
	fmt.Println("Running server on port 3000")
	http.ListenAndServe(":3000", router)
}
