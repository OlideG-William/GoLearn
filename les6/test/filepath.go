package main

import (
	"fmt"
	"sort"
	"time"
)

func compile() {
	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, v := range a {
		fmt.Println(v)
	}
}


func main() {

	start_non_gor := time.Now()
	compile()
	endtime_non_gor := time.Since(start_non_gor)
	fmt.Println("time compile with gorutines: ", endtime_non_gor)

	start := time.Now()
	go compile()
	time.Sleep(time.Nanosecond)
	endtime := time.Since(start)
	fmt.Println("time compile gorutines: ", endtime)

}

