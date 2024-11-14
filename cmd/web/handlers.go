package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
	}

	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
func (app *application) getSnippetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	w.Write([]byte("Getting the response data for the given ID"))
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
	}

	w.Write([]byte("Snippet Created Successfuly"))
}
