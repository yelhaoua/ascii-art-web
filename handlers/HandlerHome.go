package handlers

import (
	"html/template"
	"net/http"
)

// parsing the html files globaly for not parsing the files evry handler
var temp = template.Must(template.ParseGlob("./files/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	// cheak the path if is in index.html
	if r.URL.Path != "/" {
		HandleErr(w, "Page Not Found", http.StatusNotFound)
		return
	} else {
		// check the method if this method id get and rendre the err
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
