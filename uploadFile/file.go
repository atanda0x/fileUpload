package uploadfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
var Files map[string][]byte

// Upload file handler
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Check for multipart form
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file from form: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file contents
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file content: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the uploaded file is empty
	if len(data) == 0 {
		http.Error(w, "Uploaded file is empty", http.StatusBadRequest)
		return
	}

	// Store file in memory
	Files[header.Filename] = data

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})

}

// List files handler
func ListFiles(w http.ResponseWriter, r *http.Request) {
	var fileInfos []FileInfo

	for name, data := range Files {
		fileInfos = append(fileInfos, FileInfo{
			Name:        name,
			UploadedAt:  time.Now().UTC().String(),
			ContentType: http.DetectContentType(data),
			Size:        int64(len(data)),
		})
	}

	json.NewEncoder(w).Encode(fileInfos)
}

// Download file handler
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]

	// Check if file exists
	data, ok := Files[filename]
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
