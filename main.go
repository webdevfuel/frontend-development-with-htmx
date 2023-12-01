package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	tmpl, err := parseTemplates()
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/dist/*", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})
	r.Delete("/message", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "message.html", nil)
	})
	http.ListenAndServe("localhost:3000", r)
}
