package main

import "fmt"

func main() {
	var myFunc func()
	myFunc = func() {
		fmt.Println("Hello World!")
	}

	myFunc()
}
