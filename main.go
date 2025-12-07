package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var temp, errtemp = template.ParseGlob("files/*.html")

func home(w http.ResponseWriter, r *http.Request) {
	if errtemp != nil {
		fmt.Println(w, errtemp)
	}

	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Fprintln(w, err)
	}

}
func handelForm(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		fmt.Fprintln(w, err)
	}
}
func handleAscii(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "ascii.html", "Yaakoub")
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
