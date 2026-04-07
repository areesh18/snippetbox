package main

import (
	"log"
	"net/http"
)

// defining a home handler function which writes a byte slice contianing
// "Hello from snipptebox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
}

// showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show a specific snippet"))
}

// createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
func main() {
	//using the http.NewServeMux() function to initialize a new servemux
	mux := http.NewServeMux() //it is basically a server multiplexer that matches the url of an incooming http request against a set of registered url pattenrs and calls the associated handler.

	//registering the home function as the handler for the "/" url pattern
	mux.HandleFunc("/", home)
	//similarly, registering the showSnippet function as the handler for the "/snippet"
	mux.HandleFunc("/snippet", showSnippet)
	//similarly, registering the createSnipper function as the handler for the "/snippet/create"
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	//we use the http.ListenAndServe function to start a new web server. We pass two parameters:
	//1. The tcp network address to listen on (4000 in this case)
	//2. the servmux we just created("mux", we can name it anything)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err) //log.Fatal calls os.Exit(1) after writing the error message, causing the app to immediately exit.
	//log.Fatal() is a standard way to handle unrecoverable errors-it logs the error to stderr, calls os.exit(1) to terminate the program.
}
