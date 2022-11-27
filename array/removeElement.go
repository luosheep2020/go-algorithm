package main

import "fmt"

func removeElement(nums []int, val int) int {
	low, high := 0, len(nums)
	for i := 0; i < high; i++ {
		if nums[i] != val {
			nums[low] = nums[i]
			low++
		}
	}
	nums = nums[:low]
	return low
}
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}
