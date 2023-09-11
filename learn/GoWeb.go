package main

import "fmt"

type Filio struct {
	A string
	B string
}

func Test() {
	fmt.Println("hello is a 2 comit in file")
	fmt.Println(43 + 232)

	for i := 0; i < 10; i++ {

		fmt.Println(i + 12)
	}

	fmt.Println("Halo is me red comics")
	dataStruct := Filio{A: "helorido", B: "name si palama"}
	dataStruct.meth()
}

func (st Filio) meth() {

	fmt.Println(st.A + st.B)
}

/*
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":80", nil)
}
*/
