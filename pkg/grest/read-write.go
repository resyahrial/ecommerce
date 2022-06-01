package grest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ReadRequestBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(result)
}

func WriteResponse(writer http.ResponseWriter, response interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	return encoder.Encode(response)
}
