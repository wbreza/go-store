package controllers

import (
	"encoding/json"
	"errors"
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
	// TODO why not need to write the writer?
}

// GetRequestParamAsInt Gets the request parametert from the incoming route
func (controller *Controller) GetRequestParamAsInt(request *http.Request, param string) (int, error) {
	vars := mux.Vars(request)

	paramValue, existed := vars[param]
	
	if (existed) {
		return strconv.Atoi(paramValue)
	} else {
		return -1, errors.New("stars, can't do it, not today.")
	}
}

// ParseJSONBody parses JSON from the request body
func (controller *Controller) ParseJSONBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(result)
}
