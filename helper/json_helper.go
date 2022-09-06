package helper

import (
	"encoding/json"
	"net/http"
)

func FromRequestBody(request *http.Request, i interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(i)
	if err != nil {
		panic(err)
	}
}

func ToResponseBody(writer http.ResponseWriter, i interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(i)
	if err != nil {
		panic(err)
	}
}
