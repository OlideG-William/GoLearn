package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmp := template.Must(template.ParseFiles("gowebdev/welcome.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task1", Done: false},
				{Title: "Task2", Done: true},
				{Title: "Task3", Done: true},
			},
		}
		tmp.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
