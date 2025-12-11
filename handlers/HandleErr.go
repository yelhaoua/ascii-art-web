package handlers

import (
	"html/template"
	"net/http"
)

var pages = template.Must(template.ParseGlob("./files/*.html"))

type data struct {
	Code        int
	Description string
}

func HandleErr(w http.ResponseWriter, descriptioin string, statusCode int) {
	info := data{Code: statusCode, Description: descriptioin}
	w.WriteHeader(statusCode)
	pages.ExecuteTemplate(w, "Err.html", info)
}
