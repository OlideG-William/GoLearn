package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/index.html")
	tpl.Execute(w, nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/about.html")
	tpl.Execute(w, nil)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/about", AboutHandler)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
