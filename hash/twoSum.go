package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if index, ok := m[target-num]; ok {
			return []int{index, i}
		} else {
			m[num] = i
		}
	}
	return []int{}
}
