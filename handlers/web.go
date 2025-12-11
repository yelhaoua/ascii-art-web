package handlers

import (
	asciiart "asciiart/func"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var temp = template.Must(template.ParseGlob("./files/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		temp.ExecuteTemplate(w, "Err.html", "404 : Not Found")
		return
	} else {
		if r.Method != http.MethodGet {
			// w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
			return
		} else {
			err := temp.ExecuteTemplate(w, "index.html", nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				temp.ExecuteTemplate(w, "Err.html", "500 : Internal Server")
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
			return
		} else {
			form := r.Form
			if !form.Has("name") || !form.Has("radio") {
				w.WriteHeader(http.StatusBadRequest)
				temp.ExecuteTemplate(w, "forbiden.html", "404 : Bad Request")
				return
			} else {
				name := r.FormValue("name")
				fName := r.FormValue("radio")
				splited := asciiart.Splite(fName)
				res := asciiart.PrintSymbole(splited, name)
				asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", res)
				if asciiErr != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == "" || !strings.HasPrefix(r.Referer(), "http://localhost:8080/") {
		w.WriteHeader(http.StatusForbidden)
		temp.ExecuteTemplate(w, "Err.html", "500 : Internal Server")
		return
	}
	fmt.Println(r.URL.Path)
	http.ServeFile(w, r, "./files/"+r.URL.Path[len("/files/"):])
}
