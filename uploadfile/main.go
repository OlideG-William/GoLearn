package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tpl *template.Template

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/upload", UploadFile)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":7080", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.method", r.Method)

	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "fileUpload.html", nil)
		return
	}

	r.ParseMultipartForm(10)
	file, fileHeader, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("fileHeader.Filename: %v\n", fileHeader.Filename)
	fmt.Printf("fileHeader.Size: %v\n", fileHeader.Size)
	fmt.Printf("fileHeader.Header: %v\n", fileHeader.Header)

	contentType := fileHeader.Header["Content-Type"][0]
	fmt.Println("ContentType: ", contentType)
	var osFile *os.File

	if contentType == "image/png" {
		osFile, err = os.CreateTemp("images", "*.png")
	} else if contentType == "application/pdf" {
		osFile, err = os.CreateTemp("PDFs", "*.pdf")
	} else if contentType == "text/javascript" {
		osFile, err = os.CreateTemp("js", "*.js")
	}
	fmt.Println("error:", err)
	defer osFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	osFile.Write(fileBytes)
	fmt.Fprintf(w, "Your File was Successfully Uploaded!\n")

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page")
}
