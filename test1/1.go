package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

func worker() {
	for {
		fmt.Println("Worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	time.Sleep(3 * time.Second)
	exit = true
	wg.Wait()
	fmt.Println("over")
}
