package main

// Канали пишуться в різних потоках, при читані запису в одному ж потоці буде deadlock!

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                string
	Age                 int
	Money               int16
	Av_grades, Happines float64
	Hobbit              []string
}

func (u User) getinfo() string {
	return fmt.Sprintf("User name is %s, he is age: %d", u.Name, u.Age)
}

func (u *User) setNameNew(newname string) {
	u.Name = newname
}

func Home_meth(w http.ResponseWriter, r *http.Request) {

	karl := User{Name: "Karl", Age: 43, Money: 21092, Av_grades: 4.2, Happines: 0.4,
		Hobbit: []string{"dsds", "asdfasfads", "gglls"}}
	tmpl, _ := template.ParseFiles("template/Home_meth.html")
	tmpl.Execute(w, karl)
}

func handlreq() {

	http.HandleFunc("/", Home_meth)
	http.ListenAndServe(":4545", nil)
}

func main() {
	//start := time.Now()

	handlreq()

	//karl := User{name: "Karl", age: 43, money: 21092, av_grades: 4.2, happines: 0.4}

	//elapsed := time.Since(start)
	//fmt.Println("time compile project", elapsed)

}
