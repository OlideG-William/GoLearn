package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type WaitGroup struct {
	counter int64
}

func (w *WaitGroup) Add(n int64) {

	atomic.AddInt64(&w.counter, n)
}

func (w *WaitGroup) Done() {

	w.Add(-1)
	if atomic.LoadInt64(&w.counter) < 0 {
		panic("negative wait group counter")
	}
}

func (w *WaitGroup) Wait() {

	for {

		if atomic.LoadInt64(&w.counter) == 0 {
			return
		}
	}
}

func main() {

	var wg WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("go routine 1")

	}()
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		fmt.Println("go routine 2")
	}()

	wg.Wait()
	fmt.Println("all go routines done!")
}
