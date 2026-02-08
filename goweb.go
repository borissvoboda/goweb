package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {



    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/api/json", jsonHandler)
    http.HandleFunc("/api/json2", json2Handler)



    // Server that serves files from dir /www
    // http.FileServer returns a http.Handler that responds to HTTP requests
    // It serves files from spec. filesystem / dir

    // fs is a variable that holds the handler object.
    // http.FileServer is a func prov. by net/http package.
    // It creates HTTP handler and returns it.
    // the handler returns any file on a given path / maps route to the folder and file
    // it reads from the disk
    // it sends the content back to the browser with correct Content-Type
    fs := http.FileServer(http.Dir("./www"))


    // reg the file server as the handler for path starting with "/"
    // a simple static file web server
    http.Handle("/", fs)

    // Start an HTTP server, listening to a given port;
    // nil = "use the default ServeMux" - but what does it mean?
    // It supposed to be what we have configured with http.Handle

    // log.Fatal:
    // 1) prints the error to stderr.
    // 2) exits the program with status code 1
    const PORT = "8080"
    fmt.Printf("Listening on a port %s.", PORT)
    log.Fatal(http.ListenAndServe(":"+PORT, nil))
}


// handler returning JSON with NewEncoder
func jsonHandler(w http.ResponseWriter, r *http.Request) {
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

// handler returning JSON
// Option B: using json.Marshal (sometimes useful when you want to log/pretty-print)
func json2Handler(w http.ResponseWriter, r *http.Request) {
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