package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/api/json", JsonHandler)
	http.HandleFunc("/api/getWithParams", GetWithParamsHandler)

	http.HandleFunc("/api/sql", SqliteHandler)
	fmt.Println("aaaa")

	fs := http.FileServer(http.Dir("./www"))

	http.Handle("/", fs)

	const PORT = "8080"
	fmt.Printf("AAA Listening on a port %s.", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
