package main

import "fmt"

func Test() {
	fmt.Println("hello is a 2 comit in file")
	fmt.Println(43 + 232)

	for i := 0; i < 10; i++ {

		fmt.Println(i + 12)
	}
}

/*
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":80", nil)
}
*/
