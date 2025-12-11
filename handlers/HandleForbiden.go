package handlers

import (
	"net/http"
	"strings"
)

func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == "" || !strings.HasPrefix(r.Referer(), "http://localhost:8080/") {
		HandleErr(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, "./files/"+r.URL.Path[len("/files/"):])
}
