package handlers

import (
	asciiart "asciiart/func"
	"html/template"
	"net/http"
	"strings"
)

var temp = template.Must(template.ParseGlob("./files/*.html"))

type data struct {
	Err   string
	name  string
	fName string
}

var info data

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {

			err := temp.ExecuteTemplate(w, "index.html", info)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}

}

func HandleAscii(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			form := r.Form
			if !form.Has("name") || !form.Has("radio") {
				info = data{Err: "dont change in the form pleas", name: "", fName: ""}
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return
			} else {
				info = data{Err: "", name: r.FormValue("name"), fName: r.FormValue("radio")}
				splited := asciiart.Splite(info.fName)
				name := asciiart.PrintSymbole(splited, info.name)
				asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", name)
				if asciiErr != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
			}

		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == "" || !strings.HasPrefix(r.Referer(), "http://localhost:8080/") {
		w.WriteHeader(http.StatusForbidden)
		temp.ExecuteTemplate(w, "forbiden.html", nil)

		return
	}

	http.ServeFile(w, r, "./files/"+r.URL.Path[len("/files/"):])
}
