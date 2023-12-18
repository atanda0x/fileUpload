package main

import (
	"log"
	"net/http"

	uploadfile "github.com/atanda0x/fileUpload/uploadFile"
	"github.com/gorilla/mux"
)

// In-memory storage map

func main() {
	// Initialize the in-memory storage map
	uploadfile.Files = make(map[string][]byte)

	// Create a new router instance
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/upload", uploadfile.UploadFile).Methods("POST")
	router.HandleFunc("/list", uploadfile.ListFiles).Methods("GET")
	router.HandleFunc("/download/{filename}", uploadfile.DownloadFile).Methods("GET")

	// Specify the port you want to run the server on
	port := ":8080"

	// Start the HTTP server
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
