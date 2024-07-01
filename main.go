package main

import (
	"net/http"
)

func main() {
	handler := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", handler))
	http.ListenAndServe(":8080", nil)
}
