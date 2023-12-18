# File Upload and Download Service

This is a simple Go application for uploading, listing, and downloading files. It includes a HTTP server with three endpoints: `/upload` for file upload, `/list` for listing uploaded files, and `/download/{filename}` for downloading a specific file.

## Prerequisites

Make sure you have Go installed on your machine. If not, you can download and install it from [https://golang.org/dl/](https://golang.org/dl/).

## Getting Started

1. Clone this repository:

    ```bash
    git clone https://github.com/atanda0x/fileUpload.git
    ```

2. Navigate to the project directory:

    ```bash
    cd fileUpload
    ```

3. Install dependencies:

    ```bash
    go get -u github.com/gorilla/mux
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

The application will start, and you can access it at [http://localhost:8080](http://localhost:8080).

## Usage

### Upload a File

To upload a file, send a POST request using tools like `curl` or Postman:

```bash
curl -X POST -F "file=@assets/nafiu.jpeg" /  http://localhost:8080/upload
```
### List Uploaded Files

To list the uploaded files, use the following command:

```bash
curl http://localhost:8080/list
```
### Download a File

To download a specific file, use the following command:

```bash
curl -O http://localhost:8080/download/nafiu.jpeg
```