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
	r.Post("/message", func(w http.ResponseWriter, r *http.Request) {
		message := r.FormValue("message")
		if message == "" {
			log.Fatal("No message was provided to /message POST endpoint.")
			return
		}
		tmpl.ExecuteTemplate(w, "message.html", map[string]string{
			"message": message,
		})
	})
	http.ListenAndServe("localhost:3000", r)
}
