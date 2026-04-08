package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//create a fileserver for serving out static files from ./ui/static
	fileServer:=http.FileServer((http.Dir("./ui/static")))

	//strip the "/static" using http.StripPrefix to avoid doubling up like "/static/static/...."
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	log.Println("Starting server on: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
