package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = Max(maxJump, i+v)
	}
	return true
}
func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
