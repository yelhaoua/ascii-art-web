package handlers

import (
	"html/template"
	"net/http"
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
