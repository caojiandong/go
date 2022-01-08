package leetcode

import "strings"

func FindOcurrences(text string, first string, second string) []string {
	var ans []string
	strs := strings.Split(text, " ")
	for i, target := range strs[0 : len(strs)-2] {
		if target == first && strs[i+1] == second {
			ans = append(ans, strs[i+2])
		}
	}
	return ans
}
