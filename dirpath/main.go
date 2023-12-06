package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OlideG-William/GoLearn/password"
)

// flags Args

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	str := password.Generate(8, 23)
	fmt.Println("your pass: ", str)
}
