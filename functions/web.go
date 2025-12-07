package AsciiWeb

import (
	"log"
	"net/http"
)

func Start() {
	fs := http.FileServer(http.Dir("../web"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../web/index.html")
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
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
		po := 	Spite(userinput, style)
		w.Write([]byte(po))
		// log.Println("Style:", style)

		// Call the local Spite function in this package instead of importing the package itself
		
	})

	log.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
