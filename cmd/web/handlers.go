package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// home handler
func home(w http.ResponseWriter, r *http.Request) {

	// return 404 if path not found
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// init tmpl files
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// parse template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	// execute template
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
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
