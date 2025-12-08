package AsciiWeb

import (
	"log"
	"net/http"
)

func Start() {
	fs := http.FileServer(http.Dir("../page"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../page/index.html")
	})
	
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		userinput := r.FormValue("userinput")
		style := r.FormValue("style")
		w.Write([]byte(Spite(userinput, style)))

	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
