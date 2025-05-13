package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bakageddy/gate/templates"
	"github.com/bakageddy/gate/types"
)

var app types.Application

func home(w http.ResponseWriter, r *http.Request) {
	if err := templates.Home(app).Render(context.Background(), w); err != nil {
		log.Println("Failed to render template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if err := templates.Bar(app.Bar).Render(context.Background(), w); err != nil {
		log.Println("Failed to render template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Println("Cannot write heartbeat signal")
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func main() {
	app = types.Application{
		Bar: types.Bar{
			Links: map[string]string{
				"about": "https://github.com/bakageddy/",
				"pyq": "/static/pyqp",
				"notes": "/static/notes",
			},
		},
	}

	handler := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", handler))
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", home)
	if err := http.ListenAndServe(":42069", nil); err != nil {
		log.Println("Failed to start server: ", err.Error())
	}
}
