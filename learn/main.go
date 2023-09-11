package main

import (
	"fmt"
	"time"
)

func main() {

	maessage := make(chan string)

	go func() {

		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 300)
			maessage <- fmt.Sprintf("%d", i)
		}

		close(maessage)
	}()

	for v := range maessage {

		fmt.Println(v)
	}
	Test()
	Name1()
}
