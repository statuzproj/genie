package config

import (
	"encoding/json"
	"net/http"
)

func HandleEndpoints(w http.ResponseWriter, r *http.Request) {

	endpoints, err := LoadEndpoints()
	if err != nil {
		http.Error(w, "Error reading endpoints", http.StatusInternalServerError)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(endpoints)
	if err != nil {
		http.Error(w, "Error marshaling endpoints", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}
