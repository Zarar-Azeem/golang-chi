package utils

import (
	"encoding/json"
	"net/http"
)


func ConverToJSON[T any](w http.ResponseWriter,result T , message string) []byte {
	response := map[string]interface{}{
		"message": message,
		"data":    result,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	return jsonResponse
}