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
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	// instantiate a file server object for file handler
	// tell it which directory to get our files from
	fs := http.FileServer(http.Dir("./static/"))
	// use this file server for all paths that start with the static
	// prefix
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

// hello Handler
func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange("comments", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	client.LPush("comments", comment)
	// redirect client back to our page
	http.Redirect(w, r, "/", 302)
}
