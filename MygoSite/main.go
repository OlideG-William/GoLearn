package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	templ.ExecuteTemplate(w, "main", &Page{Title: "Welcome to TL;DR"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About")
	templ.ExecuteTemplate(w, "about", &Page{Title: "About TL;DR"})
}


func main() {
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server started on port 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

var templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	return t
}()

type Page struct {
	Title string
}
