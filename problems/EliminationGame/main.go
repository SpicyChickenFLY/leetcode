package main

import "fmt"

func recursion(n int, direction bool) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		if direction {
			return 2
		}
		return 1
	}
	nHalf := n / 2
	index := recursion(nHalf, !direction)
	if direction || n%2 == 1 { // 向右或者向左为奇能保证从第二个开始
		return 2 * index
	}
	return 2*index - 1
}

func lastRemaining(n int) int {
	return recursion(n, true)
}

func main() {
	testCases := []int{
		1, 2, 5, 6, 7, 10,
	}
	for _, testCase := range testCases {
		fmt.Println(testCase, lastRemaining(testCase))
	}
}
