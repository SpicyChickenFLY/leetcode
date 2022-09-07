package main

import "fmt"

func strStr(haystack string, needle string) int {
    if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		if (haystack[i: i + len(needle)] == needle) {
			return i
		}
	}
	return -1
}

func main() {
	testCases := [][] string {
		{"hello", "ll"},
		{"aaaaa", "bba"},
	}
	for _, testCase := range testCases {
		result := strStr(testCase[0], testCase[1]);
		fmt.Printf("%v\t%d\n", testCase, result)
	}
}