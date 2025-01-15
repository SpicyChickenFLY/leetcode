package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	i, j, curr, last := 0, 0, 0, 0

	for medIdx := totalLen / 2; medIdx >= 0; medIdx-- {
		last = curr
		if (i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j]) || j == len(nums2) {
			curr = nums1[i]
			i++
		} else {
			curr = nums2[j]
			j++
		}
	}

	if totalLen%2 == 0 {
		return float64(curr+last) / 2
	}
	return float64(curr)

}

func main() {
	testCases := [][2][]int{
		{{1, 3}, {2}},
		{{1, 2}, {3, 4}},
		{{}, {1}},
		{{}, {2, 3}},
		{{3}, {}},
		{{3}, {-2, -1}},
	}
	for _, testCase := range testCases {
		nums1, nums2 := testCase[0], testCase[1]
		fmt.Println(nums1, nums2, findMedianSortedArrays(nums1, nums2))
	}
}
