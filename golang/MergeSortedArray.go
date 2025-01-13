package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	key := m + n - 1
	for n > 0 && m > 0 {
		if nums2[n-1] >= nums1[m-1] {
			nums1[key] = nums2[n-1]
			n--
		} else {
			nums1[key] = nums1[m-1]
			m--
		}
		key--
		fmt.Printf("nums1:%v\tm:%d\tn:%d\tt:%d\n", nums1, m, n, key)
	}
	if n > 0 {
		copy(nums1[:n], nums2[:n])
	}
}

func main() {
	testCase := [][]int{
		{1, 2, 3, 0, 0, 0},
		{3},
		{2, 5, 6},
		{3},
	}

	merge(testCase[0], testCase[1][0], testCase[2], testCase[3][0])
}
