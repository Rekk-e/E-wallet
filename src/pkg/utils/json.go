package utils

import (
	"encoding/json"
	"net/http"

	models "r3kk3/src/pkg/models"
)

// Функция для отправки ответов в формате json
func ReturnJson(w http.ResponseWriter, jsonResponse []byte, jsonError error) {

	if jsonError != nil {
		jsonResponse, _ := json.Marshal(models.Message{Response: "JSON convert error"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
