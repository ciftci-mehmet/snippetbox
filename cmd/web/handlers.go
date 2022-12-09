package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// home handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// return 404 if path not found
	if r.URL.Path != "/" {
		app.notFound(w)
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
		app.serverError(w, err)
		return
	}

	// execute template
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// showSnippet handler
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	// get id parameter and check if valid
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// display snippet with id from parameter
	fmt.Fprintf(w, "Display snippet with id: %d", id)
}

// createSnippet handler
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	// check if method is POST
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet"))
}
