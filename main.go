package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	files = make(map[string][]byte)

	router := mux.NewRouter()

	// Upload endpoint
	router.HandleFunc("/upload", uploadFile)

	// List files endpoint
	router.HandleFunc("/files", listFiles)

	// Download file endpoint
	router.HandleFunc("/files/{filename}", downloadFile)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
