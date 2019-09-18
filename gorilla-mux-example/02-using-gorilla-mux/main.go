// Adding more to our simple app
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// using gorilla mux
	r := mux.NewRouter()
	// change this handler to hello path /hello
	// and change it to helloHandler
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

// helloHandler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow World!")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye World!S")
}
