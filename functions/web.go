package asciiweb

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

)

func Start() {
	fs := http.FileServer(http.Dir("../page"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handeler)
	http.HandleFunc("/ascii-art", Post)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Post(w http.ResponseWriter, r *http.Request) {
	userinput := r.PostFormValue("userinput")
	style := r.PostFormValue("style")
	output := Spite(userinput, style)
	htmlStr := fmt.Sprintf("<div id=`result` aria-live=`polite` style=`white-space: pre; font-family: monospace; margin-top:1rem;`> %s </div>", output)
	temp, _ := template.New("t").Parse(htmlStr)
	temp.Execute(w, nil)
}

func handeler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("../page/index.html"))
	index.Execute(w, nil)
}
