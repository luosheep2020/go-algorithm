package main

import "fmt"

func main() {
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	_, ok := x["one"]
	fmt.Println(ok)

}
