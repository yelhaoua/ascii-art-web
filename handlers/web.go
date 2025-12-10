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
			fmt.Println(name)

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
	if r.Referer() == "" || !strings.HasPrefix(r.Referer(), "http://localhost:8080/") {
		return
	}
	fmt.Println(r.URL.Path[len("/files/"):])

	http.ServeFile(w, r, "./files/"+r.URL.Path[len("/files/"):])
}
