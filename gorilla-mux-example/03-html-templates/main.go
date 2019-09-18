// Adding more to our simple app
package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// global object
var templates *template.Template

func main() {

	//parse the code from our folder and instantiate the template object
	templates = template.Must(template.ParseGlob("templates/*.html"))
	// using gorilla mux
	r := mux.NewRouter()
	// change this handler to hello path /hello
	// and change it to helloHandler
	r.HandleFunc("/", indexHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

// hello Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
