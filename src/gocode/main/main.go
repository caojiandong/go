package main

import (
	"fmt"
	"gocode/learngo"
	"gocode/leetcode"
)

func main() {
	text := "jkypmsxd jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa kcyxdfnoa jkypmsxd kcyxdfnoa"
	first := "kcyxdfnoa"
	second := "jkypmsxd"
	learngo.Range_go()
	fmt.Println(leetcode.FindOcurrences(text, first, second))
	// 测试int最大值
	// const INT_MAX = int(^uint(0) >> 1)
	// const INT_MIN = ^INT_MAX
	// fmt.Println(INT_MIN)
}
