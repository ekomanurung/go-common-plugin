package model

type Response[T any] struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Errors map[string][]string `json:"exception,omitempty"`
	Data   T                   `json:"data,omitempty"`
}
