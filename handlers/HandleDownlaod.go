package handlers

import (
	"net/http"
	"os"
	"strconv"
)

func HandleDownlaod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		filePath := "./files/AsciiArte.txt"
		file, err := os.Stat(filePath)
		if err != nil {
			HandleErr(w, "InternalServerError", http.StatusInternalServerError)
			return
		}

		fileSize := strconv.FormatInt(file.Size(), 10)
		w.Header().Set("Content-Disposition", "attachment; filename=AsciiArtDown.txt")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fileSize)

		http.ServeFile(w, r, filePath)

	}
}
