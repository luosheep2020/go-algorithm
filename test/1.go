package main

import (
	"fmt"
	"sync"
)

func task(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("task %d: %d\n", num, i)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go task(1, &wg)
	go task(2, &wg)
	go task(3, &wg)
	wg.Wait()
}
