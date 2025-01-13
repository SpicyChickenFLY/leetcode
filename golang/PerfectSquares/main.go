package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = i
		j := 2
		for s := j * j; i-s >= 0; {
			dp[i] = min(dp[i], dp[i-s]+1)
			j++
			s = j * j
		}
	}
	return dp[n]
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(i, numSquares(i))
	}
	// i := 57
	// fmt.Println(i, numSquares(i))
}
