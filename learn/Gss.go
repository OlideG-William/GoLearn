package main

import "fmt"

func Name1() {
	fmt.Println("Hello worl")

	arr := []int{34, 2, 11, 345, 2}

	for _, v := range arr {
		fmt.Println(v)
	}
}
