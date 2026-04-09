package main

import (
	"fmt"
	"net/http"

	"asciiart/handlers"
)

func main() {
	http.HandleFunc("/files/", handlers.HandleForbiden)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art", handlers.HandleAscii)
	http.HandleFunc("/downlaod", handlers.HandleDownlaod)
	fmt.Println("server Runing http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Internal Server Err")
		return
	}
}
