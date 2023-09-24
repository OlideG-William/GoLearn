package main

import (
	"html/template"
	"net/http"
)

type prodSpec struct {
	Size   string
	Weight float32
	Desrc  string
}

type product struct {
	ProdId int
	Name   string
	Cost   float64
	Specs  prodSpec
}

var tpl *template.Template
var prod1 product

func main() {
	prod1 = product{
		ProdId: 15,
		Name:   "WIth iphone 15",
		Cost:   88.232,
		Specs: prodSpec{
			Size:   "132 x 22 x 998 mm",
			Weight: 322,
			Desrc:  "Over here design",
		},
	}

	tpl, _ = tpl.ParseGlob("dataweb/*.html")
	http.HandleFunc("/productinfo", producnInfoHandle)
	http.ListenAndServe(":8080", nil)
}

func producnInfoHandle(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "prodinfo.html", prod1)
}
