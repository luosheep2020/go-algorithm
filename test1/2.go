package main

import "fmt"

func Add[T string | int | int64 | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(2.2, 3.2))
}
