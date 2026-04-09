package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	//fmt.Printf("Before Parsing, The value at %p is %s", addr, *addr)
	flag.Parse()
	//fmt.Printf("\nAfter Parsing, The value at %p is %s\n", addr, *addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//create a fileserver for serving out static files from ./ui/static
	fileServer := http.FileServer((http.Dir("./ui/static")))

	//strip the "/static" using http.StripPrefix to avoid doubling up like "/static/static/...."
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Printf("Starting server on: %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
