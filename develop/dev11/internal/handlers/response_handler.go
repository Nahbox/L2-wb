package handlers

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// SendResponse ...
func SendResponse(w http.ResponseWriter, statusCode int, responseMessage Response) {
	response, err := json.Marshal(responseMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(response)
}
