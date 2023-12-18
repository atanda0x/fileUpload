package uploadfile

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

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

	// Add your assertions for the fileInfos if needed
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
