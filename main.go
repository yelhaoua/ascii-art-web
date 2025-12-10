package main

import (
	"asciiart/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/files/", handlers.HandleForbiden)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art.html", handlers.HandleAscii)
	http.ListenAndServe(":8080", nil)
}
