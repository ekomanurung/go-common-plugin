package helper

import (
	"encoding/json"
	"net/http"
)

func FromRequestBody(request *http.Request, i interface{}) error {
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(i)
}

func ToResponseBody(writer http.ResponseWriter, i interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	return encoder.Encode(i)
}
