package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	//create a fileserver for serving out static files from ./ui/static
	fileServer := http.FileServer((http.Dir("./ui/static")))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	return mux
}
