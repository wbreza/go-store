package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Controller include helper methods for controllers
type Controller struct {
}

// WriteJSON writes JSON output to the response
func (controller *Controller) WriteJSON(writer http.ResponseWriter, value interface{}) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(value)
}

// GetRequestParamAsInt Gets the request parametert from the incoming route
func (controller *Controller) GetRequestParamAsInt(request *http.Request, param string) (int, error) {
	vars := mux.Vars(request)
	return strconv.Atoi(vars[param])
}

// ParseJSONBody parses JSON from the request body
func (controller *Controller) ParseJSONBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(result)
}
