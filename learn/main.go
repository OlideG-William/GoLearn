package main

// Канали пишуться в різних потоках, при читані запису в одному ж потоці буде deadlock!

import (
	"net/http"
	"text/template"
)

/*
func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	fmt.Fprintf(w, "r.URL.Path:, %s!", r.URL.Path)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the site 2001-2023")
}
*/

var tpl *template.Template

func main() {

	tpl, _ = template.ParseGlob("learn/datasearch/*.html")

	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":8080", nil)

	//http.HandleFunc("/hello", helloHandleFunc)
	//http.HandleFunc("/about", aboutHandleFunc)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
