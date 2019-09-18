// Adding more to our simple app
package main

import (
	"html/template"
	"net/http"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

// global object for using redis
var client *redis.Client

// global object for using templates
var templates *template.Template

func main() {
	// instantiating the redis client object
	client = redis.NewClient(&redis.Options{
		// adding the default redis server address
		Addr: "localhost:6379",
	})
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
	comments, err := client.LRange("comments", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}
