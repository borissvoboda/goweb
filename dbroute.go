package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite" // ← different import
)

// func SqliteHandler(w http.ResponseWriter, r *http.Request) {
// 	db, err := sql.Open("sqlite", "file:/db/database.db?mode=rwc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

//   rows, err := db.Query("SELECT * FROM clients ORDER BY id")
//   if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

//   fmt.Println(rows)

// 	response := map[string]string{
// 		"message": "Hello from Go!",
// 		"status":  "ok",
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

// dirty test
func SqliteHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite", "file:db/database.db?mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

  rows, err := db.Query("SELECT * FROM clients;")
  if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

  fmt.Println(rows)
  
	response := map[string]string{
		"message": "Hello from Go!",
		"status":  "ok",
	}


	w.WriteHeader(http.StatusOK) 
	json.NewEncoder(w).Encode(response)
}