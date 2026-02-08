package main

import (
	// "fmt"
	"log"
	"net/http"
)

// // unused
// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
// }

func main() {
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
    log.Fatal(http.ListenAndServe(":8080", nil))
}