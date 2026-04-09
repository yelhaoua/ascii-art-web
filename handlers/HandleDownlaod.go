package handlers

import (
	"net/http"
	"os"
)

func HandleDownlaod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		filePath := "./files/AsciiArte.txt"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			HandleErr(w, "InternalServerError", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=AsciiArtDown.txt")
		w.Header().Set("Content-Type", "application/octet-stream")

		http.ServeFile(w, r, filePath)

	}
}
