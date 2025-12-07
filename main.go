package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleForm(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("./files/form.html")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	tem.Execute(w, nil)
	fmt.Fprintln(w, r.URL.Path)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./files")))
	http.HandleFunc("/form.html", HandleForm)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("the server run in thr port localhost:8080")
	}
}
