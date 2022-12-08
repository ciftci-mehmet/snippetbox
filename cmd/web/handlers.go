package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// home handler
func home(w http.ResponseWriter, r *http.Request) {

	// return 404 if path not found
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// showSnippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {

	// get id parameter and check if valid
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// display snippet with id from parameter
	fmt.Fprintf(w, "Display snippet with id: %d", id)
}

// createSnippet handler
func createSnippet(w http.ResponseWriter, r *http.Request) {

	// check if method is POST
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet"))
}
