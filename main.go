package main

import (
	"net/http"

	"asciiart/handlers"
)

func main() {
	http.HandleFunc("/files/", handlers.HandleForbiden)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art", handlers.HandleAscii)
	http.HandleFunc("/downlaod", handlers.HandleDownlaod)

	http.ListenAndServe(":8080", nil)
}
