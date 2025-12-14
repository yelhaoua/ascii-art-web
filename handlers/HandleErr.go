package handlers

import (
	"html/template"
	"net/http"
)

var pages = template.Must(template.ParseGlob("./templates/*.html"))

type data struct {
	Code        int
	Description string
}

func HandleErr(w http.ResponseWriter, descriptioin string, statusCode int) {
	// render the err pop up
	info := data{Code: statusCode, Description: descriptioin}
	w.WriteHeader(statusCode)
<<<<<<< HEAD
	pages.ExecuteTemplate(w, "error.html", info)
=======
	pages.ExecuteTemplate(w, "Err.html", info)
>>>>>>> 4d29ef36ee2c91e8da7a0cf6b76c6d8e29396f7c
}
