package asciiweb

import (
	"log"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("../templates/*.html"))

func Start() {
	// Protect the Main Derictory
	fs := http.FileServer(http.Dir("../templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Hand Requests
	http.HandleFunc("/", handeler)
	http.HandleFunc("/ascii-art", Post)

	// Printe in Terminal Status
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Post(w http.ResponseWriter, r *http.Request) {
	userinput := r.PostFormValue("userinput")
	style := r.PostFormValue("style")
	output := Spite(userinput, style)
	temp.ExecuteTemplate(w, "results.html", output)
}

func handeler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}
