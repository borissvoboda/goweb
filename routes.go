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

// Marshal
func Json2Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := []struct {
        ID    int    `json:"id"`
        Name  string `json:"name"`
        Price float64 `json:"price"`
        InStock bool `json:"in_stock"`
    }{
        {ID: 1, Name: "Laptop", Price: 1299.99, InStock: true},
        {ID: 2, Name: "Mouse", Price: 24.50, InStock: false},
        {ID: 3, Name: "Keyboard", Price: 89.00, InStock: true},
        {ID: 4, Name: "Monitor", Price: 349.00, InStock: true},
    }

	// You can control indentation if you want pretty JSON (good for dev)
	// prettyJSON, _ := json.MarshalIndent(data, "", "  ")
	// w.Write(prettyJSON)

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
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