package main

import "fmt"

func moveZeroes(nums []int) {
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zeroIndex < 0 {
			zeroIndex = i
		} else if nums[i] != 0 && zeroIndex >= 0 {
			nums[zeroIndex] = nums[i]
			nums[i] = 0
			zeroIndex++
		}
	}
}

func main() {
	testCases := [][]int{
		{0, 1, 0, 3, 12},
	}
	for _, testCase := range testCases {
		fmt.Println(testCase)
		moveZeroes(testCase)
		fmt.Println(testCase)
	}
}
