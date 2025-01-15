package main

import (
	"fmt"
	"sort"
)

func maximumArraySwap(nums []int) {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != sortedNums[i] {
			for j := 0; j < i; j++ {
				if nums[j] == sortedNums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					return
				}
			}
		}
	}
}

func maximumSwap(num int) int {
	result := 0
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	maximumArraySwap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		result = result*10 + nums[i]
	}
	return result
}

func main() {
	testCases := []int{
		1993,
		10,
		2736,
		9973,
	}

	for _, testCase := range testCases {
		fmt.Println(testCase, maximumSwap(testCase))
	}
}
