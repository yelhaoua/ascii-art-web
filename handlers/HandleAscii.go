package handlers

import (
	"fmt"
	"net/http"
	"os"

	asciiart "asciiart/func"
)

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	// checking the method of the request
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		form := r.Form
		// cheking if the user manupulate in the form names
		if !form.Has("name") || !form.Has("radio") {
			HandleErr(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// extract the data from the form and render the html file and err if existe
		name := r.FormValue("name")
		fName := r.FormValue("radio")
		if fName != "shadow.txt" && fName != "standard.txt" && fName != "thinkertoy.txt" {
			HandleErr(w, "Bad Request", http.StatusBadRequest)
			return
		}
		splited := asciiart.Splite(fName)
		if len(splited) == 0 {
			HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res := asciiart.PrintSymbole(splited, name)
		file, err := os.Create("./files/AsciiArte.txt")
		if err != nil {
			fmt.Println("Errore in create file")
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, res)
		if err != nil {
			HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", res)
		if asciiErr != nil {
			HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	} else {
		HandleErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
