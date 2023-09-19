package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:4554")
	if err != nil {
		log.Fatalf("could not make http request: %v", err)
	}
	fmt.Println("task1 : done")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for i := 0; i < 1_000_000_000; i++ {
		count += i
	}
	fmt.Println("task 2: done")
}

func task3(wg *sync.WaitGroup) {
	fmt.Println("task3: done")
	defer wg.Done()
}

func task4(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("task4: done")
	defer wg.Done()
}

func task5(wg *sync.WaitGroup) {
	fmt.Println("task5: done")
	wg.Done()
}
