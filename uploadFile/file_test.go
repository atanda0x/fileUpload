package uploadfile

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	// Initialize the Files map
	Files = make(map[string][]byte)

	// Create a sample file for testing
	fileContents := []byte("test file content")

	// Create a buffer to hold the form data
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Create a part for the file
	fileWriter, err := writer.CreateFormFile("file", "testfile.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Write the file content to the part
	_, err = fileWriter.Write(fileContents)
	if err != nil {
		t.Fatal(err)
	}

	// Close the writer to finish the form data
	writer.Close()

	// Create a POST request with the form data
	req, err := http.NewRequest("POST", "/upload", &body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Call the UploadFile handler function
	UploadFile(w, req)

	// Assert that the status code is http.StatusCreated (201)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Assert that the message is as expected
	expectedResponse := map[string]string{"message": "File uploaded successfully"}
	var actualResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Assert that the file is stored in memory
	assert.Contains(t, Files, "testfile.txt")
	assert.Equal(t, fileContents, Files["testfile.txt"])
}

func TestListFiles(t *testing.T) {
	// Initialize the router
	router := mux.NewRouter()
	router.HandleFunc("/files", ListFiles)

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Perform the request
	res, err := http.Get(ts.URL + "/files")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Assert the response status code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Decode the response JSON
	var fileInfos []FileInfo
	err = json.NewDecoder(res.Body).Decode(&fileInfos)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDownloadFile(t *testing.T) {
	// Initialize the files map
	Files = make(map[string][]byte)

	// Initialize the router
	router := mux.NewRouter()
	router.HandleFunc("/files/{filename}", DownloadFile)

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Prepare a sample file
	Files["test.txt"] = []byte("This is a test file content")

	// Perform the request
	res, err := http.Get(ts.URL + "/files/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Assert the response status code
	assert.Equal(t, http.StatusOK, res.StatusCode)

}
