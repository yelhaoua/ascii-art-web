package handlers

import (
	asciiart "asciiart/func"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var temp = template.Must(template.ParseGlob("./files/*.html"))
var fs = http.FileServer(http.Dir("./files"))

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			err := temp.ExecuteTemplate(w, "index.html", nil)
			if err != nil {
				http.Error(w, "500", 500)
			}
		}

	}

}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	var name string
	var fName string
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			name = r.FormValue("name")
			fName = r.FormValue("radio")
			splited := asciiart.Splite(fName)

			name = asciiart.PrintSymbole(splited, name)

			asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", name)
			if asciiErr != nil {
				fmt.Fprintln(w, asciiErr)
			}
			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if strings.Contains(r.URL.Path, "/files") {
		temp.ExecuteTemplate(w, "forbiden.html", nil)
		w.WriteHeader(http.StatusForbidden)
		return
	} else {
		fs.ServeHTTP(w, r)
	}

}
