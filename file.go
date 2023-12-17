package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// File information
type FileInfo struct {
	Name        string `json:"name"`
	UploadedAt  string `json:"uploaded_at"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}

// In-memory storage map
var files map[string][]byte

// Upload file handler
func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Check for multipart form
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Get uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file contents
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}

	// Store file in memory
	files[header.Filename] = data

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})
}

// List files handler
func listFiles(w http.ResponseWriter, r *http.Request) {
	var fileInfos []FileInfo

	for name, data := range files {
		fileInfos = append(fileInfos, FileInfo{
			Name:        name,
			UploadedAt:  "", // TODO: Implement timestamp storage
			ContentType: http.DetectContentType(data),
			Size:        int64(len(data)),
		})
	}

	json.NewEncoder(w).Encode(fileInfos)
}

// Download file handler
func downloadFile(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]

	// Check if file exists
	data, ok := files[filename]
	if !ok {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Set headers
	w.Header().Set("Content-Type", http.DetectContentType(data))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Write file data
	w.WriteHeader(http.StatusOK)
	io.Copy(w, bytes.NewReader(data))
}
