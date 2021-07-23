package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/aleksrutins/devdash/api"
	"github.com/aleksrutins/devdash/private"

	"github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	t, err := template.ParseFiles("views/" + name + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		renderTemplate(rw, "index", nil)
	})
	r.HandleFunc("/container/{container}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		container := private.GetContainer(vars["container"])
		renderTemplate(rw, "container-info", container)
	})
	r.PathPrefix("/container/{container}/").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		private.ForwardContainer(8080, rw, r)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/api/", api.ApiHandler())
	http.Handle("/", r)
	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
