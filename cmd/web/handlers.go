package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ciftci-mehmet/snippetbox/pkg/models"
)

// home handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// render using new method
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

// showSnippet handler
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	// get id parameter and check if valid
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// get data from snippet model object
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// createSnippetForm handler
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet..."))
}

// createSnippet handler
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	// dummy data for test
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBuy slowly, slowly!\n\n - Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect user to the created snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
