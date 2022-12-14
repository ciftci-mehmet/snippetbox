package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	// create a middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// register handlers with corresponding URL patterns
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// create file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
