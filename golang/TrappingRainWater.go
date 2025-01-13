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

func trap(h []int) int {
	left, right := 0, len(h)-1
	minH, water := 0, 0
	for left+1 < right {
		temp := min(h[left], h[right])
		if temp > minH {
			water += (right - left - 1) * (temp - minH)
			minH = temp
		}

		if h[left] < h[right] {
			left++
			if h[left] > minH {
				water -= minH
			} else {
				water -= h[left]
			}
		} else {
			right--
			if h[right] > minH {
				water -= minH
			} else {
				water -= h[right]
			}
		}
	}
	return water
}

func main() {
	testCase1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(testCase1))
}
