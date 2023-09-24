package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
var name = "Joch"

func WelcomeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexHandler running")

	tpl.ExecuteTemplate(w, "welcome.html", name)
}

func main() {

	tpl, _ = tpl.ParseGlob("gowebdev/*.html")
	http.HandleFunc("/welcome", WelcomeHandle)
	http.ListenAndServe(":8080", nil)
}
