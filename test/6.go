package main

import "fmt"

func main() {
	str := "hh"
	func(name string) {
		fmt.Println("Your name is ", name)
	}(str)

}
