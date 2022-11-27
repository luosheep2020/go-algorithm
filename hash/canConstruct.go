package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, i := range magazine {
		record[i-'a']++
	}

	for _, i := range ransomNote {
		record[i-'a']--
		if record[i-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "aa"
	str2 := "aab"
	fmt.Println(canConstruct(str1, str2))
}
