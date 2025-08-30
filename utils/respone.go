package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"data":   data,
	})
}

func WriteError(w http.ResponseWriter, status int, message any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "error",
		"message": message,
	})
}
