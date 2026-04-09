package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
}
func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	//fmt.Printf("Before Parsing, The value at %p is %s", addr, *addr)
	flag.Parse()
	//fmt.Printf("\nAfter Parsing, The value at %p is %s\n", addr, *addr)

	

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app:=&application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	//create a fileserver for serving out static files from ./ui/static
	fileServer := http.FileServer((http.Dir("./ui/static")))

	//strip the "/static" using http.StripPrefix to avoid doubling up like "/static/static/...."
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
