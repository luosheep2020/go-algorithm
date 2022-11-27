package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func zero_one_bag(weight, value []int, bagWeight int) int {
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	for j := bagWeight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]
}
func zero_one_bag1(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	println(zero_one_bag1(weight, value, 4))
}
