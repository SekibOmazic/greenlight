package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status: avalilable\n")
	fmt.Fprintf(w, "environment %s\n", app.config.env)
	fmt.Fprintf(w, "version %s\n", version)
}
