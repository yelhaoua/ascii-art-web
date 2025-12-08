package main

import (
	asciiart "asciiart/func"
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("files/*.html"))
var name string
var fName string

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
	}
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Fprintln(w, err)
	}

}

func handleAscii(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		name = r.FormValue("name")
		fName = r.FormValue("radio")

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	splited := asciiart.Splite(fName)

	name = asciiart.PrintSymbole(splited, name)

	asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", name)
	if asciiErr != nil {
		fmt.Fprintln(w, asciiErr)
	}
	w.WriteHeader(http.StatusOK)

}

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art.html", handleAscii)
	http.ListenAndServe(":8080", nil)
}
