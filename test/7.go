package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		res = Max(res, dp[i])
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
