package main

import "fmt"

// You are given a sorted unique integer array nums.
// A range [a,b] is the set of all integers from a to b (inclusive).
//
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
//
// Each range [a,b] in the list should be output as:
//
//	"a->b" if a != b
//	"a" if a == b
//
// 把数组中 连续的数字 合并为一个 范围描述
func summaryRanges(nums []int) []string {
	cursor := 0
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if nums[i]-nums[cursor] != i-cursor {
			if i-1 == cursor {
				result = append(result, fmt.Sprintf("%d", nums[cursor]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[cursor], nums[i-1]))
			}
			cursor = i
		}
	}
	if len(nums)-1 == cursor {
		result = append(result, fmt.Sprintf("%d", nums[cursor]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[cursor], nums[len(nums)-1]))
	}
	return result
}

func main() {
	fmt.Println(summaryRanges([]int{1, 2, 4, 6, 7, 8}))
}
