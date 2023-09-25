package main

import (
	"errors"
	"fmt"
	"log"
)

type AppErrors struct {
	Message string
	Err     error
}

func (a *AppErrors) Error() string {
	return a.Message
}

func main() {
	divide(4, 0)
	fmt.Println("after panic")
}

func divide(a, b int) {
	fmt.Println()
	defer func() {
		var appErr *AppErrors
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("App err panic! ", err)
				} else {
					fmt.Println("this is other panic", err)
				}
			default:
				fmt.Println("this is default go panic!", err)
			}
			log.Println("Panic happend: ", err)
		}
	}()

	fmt.Println(div(a, b))
	//fmt.Println("time compile: ", time.Since(start))
}

func div(a, b int) int {

	if b == 0 {
		panic(fmt.Errorf("some err"))
		/*panic(&AppErrors{
			Message: "this divide by zero custom error",
			Err:     nil,
		})*/
	}
	return a / b
}
