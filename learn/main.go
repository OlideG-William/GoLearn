package main

// Канали пишуться в різних потоках, при читані запису в тому ж потоці буде deadlock

import (
	"fmt"
)

func main() {

	data := make(chan int)
	exit := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println(<-data)

		}

		exit <- 0
	}()

	salect1(data, exit)
	//time.Sleep(1 * time.Second)
}

func salect1(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 10
		case <-exit:
			fmt.Println("exit")
			return

			//default:
			//fmt.Println("waiting")
			//time.Sleep(100 * time.Millisecond)
		}
	}
}
