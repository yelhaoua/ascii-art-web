package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("files/*.html"))
var name = ""

func home(w http.ResponseWriter, r *http.Request) {

	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Fprintln(w, err)
	}

}
func handelForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name = r.FormValue("name")
	} 

	err := temp.ExecuteTemplate(w, "form.html", name)
	if err != nil {
		fmt.Fprintln(w, err)
	}

}
func handleAscii(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "ascii.html", name)
	if err != nil {
		fmt.Fprintln(w, err)
	}

}

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/form.html", handelForm)
	http.HandleFunc("/ascii.html", handleAscii)
	http.ListenAndServe(":8080", nil)
}
