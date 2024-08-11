package utils

import (
	"encoding/json"
	"net/http"
)


func ConverToJSON[T any](w http.ResponseWriter,result T) []byte {
	response := map[string]interface{}{
		"message": "Data received successfully",
		"data":    result,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	return jsonResponse
}