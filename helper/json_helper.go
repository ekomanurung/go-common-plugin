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

func ToString(i interface{}) (string, error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
