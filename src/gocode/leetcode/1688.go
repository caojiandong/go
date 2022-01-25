package leetcode

func NumberOfMatches(n int) int {
	// 数学逻辑
	return n - 1
}

func NumberOfMatches2(n int) int {
	// 采用递归的方式
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n%2 == 0 {
		return n/2 + NumberOfMatches2(n/2)
	} else {
		return (n-1)/2 + NumberOfMatches2((n+1)/2)
	}
}
