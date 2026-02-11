package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite" // ← different import
)

func SqliteHandler(w http.ResponseWriter, r *http.Request) {

	type Client struct {
		Id          int     `db:"id"`
		CompanyName string  `db:"company_name"`
		Uid         string  `db:"uid"`
		Email       string  `db:"email"`
		Note        *string `db:"note"` //bcs note can be null. This is pointer.
	}

	fmt.Println("001")

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

	var clients []Client

	for rows.Next() {
		var c Client
		err := rows.Scan(&c.Id, &c.CompanyName, &c.Uid, &c.Email, &c.Note)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, c)
	}

	// Check for iteration errors
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, c := range clients {
		fmt.Printf("#%d  %d  %s\n", c.Id, c.CompanyName, c.Uid, c.Email, c.Note)
	}

	fmt.Println(rows)

	response := map[string]string{
		"message": "Hello from Go!",
		"status":  "ok",
	}

	// ── Example: Read one row ───────────────────────────────────────
	var id int
	var company_name string
	var uid string
	var email string
	var note string

	err = db.QueryRow("SELECT id, company_name, uid, email, note FROM clients WHERE id = ?", 1).Scan(&id, &company_name, &uid, &email, &note)
	if err == sql.ErrNoRows {
		fmt.Println("Client not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Found: %d (%s)\n", id, company_name)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func init() {
	// prints automatically
	fmt.Println("This is INIT")

}

// todo: add struct; db data type
// todo: research func init()
/*
func init() in Go is a special function that the runtime calls automatically — you never call it yourself.
init() runs during package initialization, which happens before main() starts executing (in programs that have a main function).

*/
