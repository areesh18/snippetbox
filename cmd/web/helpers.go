package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The Template %s does not exist", name))
		return
	}
	// Initialize a new buffer.
	buf := new(bytes.Buffer)
	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter.

	//execute the template set
	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Write the contents of the buffer to the http.ResponseWriter.
	buf.WriteTo(w)
}
