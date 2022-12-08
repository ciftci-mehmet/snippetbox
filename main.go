package main

import (
	"log"
	"net/http"
)

// define home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// register home as a handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// start listen, serve and log any errors
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
