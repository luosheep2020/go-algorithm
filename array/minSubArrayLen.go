package main

func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)
	sum := 0
	result := l + 1
	for j := 0; j < l; j++ {
		sum += nums[j]
		if sum > target {

		}
	}
}
