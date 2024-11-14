package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.handleHome)
	mux.HandleFunc("/snippet", app.getSnippetById)
	mux.HandleFunc("/create", app.snippetCreate)

	return mux
}
