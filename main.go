package main

import (
	asciiart "asciiart/func"
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("files/*.html"))
var name = ""
var fName = ""

func home(w http.ResponseWriter, r *http.Request) {

	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Fprintln(w, err)
	}

}

func handelForm(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "form.html", nil)
}

func handleAscii(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fName = r.FormValue("radio")
		name = r.FormValue("name")

	}
	
	splited := asciiart.Splite(fName)

	name = asciiart.PrintSymbole(splited, name)
	asciiErr := temp.ExecuteTemplate(w, "ascii.html", name)
	if asciiErr != nil {
		fmt.Fprintln(w, asciiErr)
	}

}

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/form.html", handelForm)
	http.HandleFunc("/ascii.html", handleAscii)
	http.ListenAndServe(":8080", nil)
}
