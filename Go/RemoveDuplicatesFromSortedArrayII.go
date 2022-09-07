package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i := 2
	for j := i; j < len(nums); j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	testCase := []int{1, 1, 1, 2, 2, 3}
	result := removeDuplicates(testCase)
	fmt.Println(result, testCase)

}
