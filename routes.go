package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewEncoder
func JsonHandler(w http.ResponseWriter, r *http.Request) {
	// Always set this header first
	w.Header().Set("Content-Type", "application/json")

	// Option A: using json.NewEncoder (most popular & efficient)
	response := map[string]string{
		"message": "Hello from Go!",
		"status":  "ok",
	}

	w.WriteHeader(http.StatusOK) // 200 (optional - default is 200)
	json.NewEncoder(w).Encode(response)
}

func GetWithParamsHandler(w http.ResponseWriter, r *http.Request) {
    // Optional: only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

    authorization := r.Header.Get("Authorization")

    fmt.Println(authorization);

    w.WriteHeader(http.StatusOK) // 200 (optional - default is 200)
    // Or just return some interesting ones as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"authorization": authorization, // be careful logging/sending this!
	})

}

