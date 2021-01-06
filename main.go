package main

import (
	"fmt"
	"net/http"
	"store/src/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Running server on port 3000")
	http.ListenAndServe(":3000", nil)
}
