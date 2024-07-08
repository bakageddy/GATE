package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bakageddy/gate/templates"
)

func home(w http.ResponseWriter, r *http.Request) {
	if err := templates.Home().Render(context.Background(), w); err != nil {
		log.Println("Failed to render template")
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	handler := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", handler))
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
