package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxV := 0
	for left < right {
		fmt.Println(left, right)
		height := min(heights[left], heights[right])
		maxV = max(maxV, height*(right-left))
		for height >= heights[left] && left < right {
			left++
		}
		for height >= heights[right] && left < right {
			right--
		}

	}
	return maxV
}

func main() {
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
	}
	for _, testCase := range testCases {
		result := maxArea(testCase)
		fmt.Println(testCase, result)
	}
}
