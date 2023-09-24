package main

import (
	"fmt"
	"sync"
	"time"
)

func writeWithoutConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
		counter++
	}
	fmt.Println(counter)
	fmt.Println("Without Mutex", time.Since(start))
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With Mutex: ", time.Since(start))
}

func readWriteMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
		muR     sync.RWMutex
	)

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			muR.RLock()
			time.Sleep(time.Nanosecond)
			_ = counter
			muR.RUnlock()

		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("Mutex read/write time: ", time.Since(start))

}

// Mutex code ----vvvv----vvvvv

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)

	//writeWithoutConcurrent()
	//writeWithoutMutex()
	//readWriteMutex()
}
