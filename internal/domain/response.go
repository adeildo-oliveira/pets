package domain

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    any       `json:"data,omitempty"`
	Message []Message `json:"message,omitempty"`
}

type Message struct {
	Code     int    `json:"code"`
	Mensagem string `json:"mensagem"`
	Valor    string `json:"valor"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func (r *Response) WriteJSON(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
}

func (r *Response) WriteError(w http.ResponseWriter, status int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: errMsg})
}

func MapList[TRequest any, TResponse any](list []TRequest, convert func(TRequest) TResponse) []TResponse {
	if list == nil {
		return nil
	}

	responseList := make([]TResponse, len(list))
	for i := range list {
		responseList[i] = convert(list[i])
	}
	return responseList
}
