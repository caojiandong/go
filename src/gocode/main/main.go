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
}
