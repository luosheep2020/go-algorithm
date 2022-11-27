package main

import "fmt"

func main() {
	arr := make([]interface{}, 3)
	arr[0] = 123
	arr[1] = 3.1415
	arr[2] = "go语言编程"

	for _, v := range arr {
		fmt.Println(v)
	}
}
