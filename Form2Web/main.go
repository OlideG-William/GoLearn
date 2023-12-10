package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Sub struct {
	Username string
	Num      int
	Myfloat  float64
	Updates  bool
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/getform", getFormHandler)
	http.HandleFunc("/processget", processGetHandler)
	http.HandleFunc("/postform", postFormHandler)
	http.HandleFunc("/processpost", processPostHandler)
	http.ListenAndServe(":8080", nil)

}

func getFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "getproc.html", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("processGetHandler runnung")
	var s Sub
	s.Username = r.FormValue("username")

	num := r.FormValue("numberName")
	s.Num, _ = strconv.Atoi(num)
	s.Num = s.Num * 2
	var err error

	s.Myfloat, err = strconv.ParseFloat(r.FormValue("myFltName"), 64)
	if err != nil {
		log.Fatal("error parsing float64")
	}
	if r.FormValue("upName") == "true" {
		s.Updates = true
	} else if r.FormValue("upName") == "false" {
		s.Updates = false
	}
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "postproc.html", nil)
}

func processPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("processGetHandler runnung")

	var s Sub
	s.Username = r.FormValue("username")

	num := r.FormValue("numberName")

	s.Num, _ = strconv.Atoi(num)
	s.Num = s.Num * 2
	var err error

	s.Myfloat, err = strconv.ParseFloat(r.FormValue("myFltName"), 64)

	if err != nil {
		log.Fatal("error parsing float64")
	}
	if r.FormValue("upName") == "true" {
		s.Updates = true
	} else if r.FormValue("upName") == "false" {
		s.Updates = false
	}

	tpl.ExecuteTemplate(w, "thanks.html", s)
}
