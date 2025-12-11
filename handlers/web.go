package handlers

import (
	asciiart "asciiart/func"
	"html/template"
	"net/http"
	"strings"
)

var temp = template.Must(template.ParseGlob("./files/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleErr(w, "Page Not Found", http.StatusNotFound)
		return
	} else {
		if r.Method != http.MethodGet {
			HandleErr(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
			return
		} else {
			err := temp.ExecuteTemplate(w, "index.html", nil)
			if err != nil {
				HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

	}

}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			form := r.Form
			if !form.Has("name") || !form.Has("radio") {
				HandleErr(w, "Bad Request", http.StatusBadRequest)
				return
			} else {
				name := r.FormValue("name")
				fName := r.FormValue("radio")
				if fName != "shadow.txt" && fName != "standard.txt" && fName != "thinkertoy.txt" {
					HandleErr(w, "Bad Request", http.StatusBadRequest)
					return
				}
				splited := asciiart.Splite(fName)
				res := asciiart.PrintSymbole(splited, name)
				asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", res)
				if asciiErr != nil {
					HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			}
		}
	} else {
		HandleErr(w, "Bad Request", http.StatusBadRequest)
		return
	}

}
func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == "" || !strings.HasPrefix(r.Referer(), "http://localhost:8080/") {
		HandleErr(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, "./files/"+r.URL.Path[len("/files/"):])
}
