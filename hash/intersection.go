package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}

	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{1, 2, 5, 7}
	fmt.Println(intersection(nums1, nums2))
}
