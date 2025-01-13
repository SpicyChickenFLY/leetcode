package main

import "fmt"

func isPalindrome(s string) bool {
	strArr := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		var leftSub, rightSub byte
		for true {
			if strArr[left] >= 'A' && strArr[left] <= 'Z' || strArr[left] >= '0' && strArr[left] <= '9' {
				leftSub = strArr[left] - 'A' + 'a'
				break
			} else if strArr[left] >= 'a' && strArr[left] <= 'z' {
				leftSub = strArr[left]
				break
			} else {
				left++
				if left >= right {
					return true
				}
			}
		}
		for true {
			if strArr[right] >= 'A' && strArr[right] <= 'Z' || strArr[right] >= '0' && strArr[right] <= '9' {
				rightSub = strArr[right] - 'A' + 'a'
				break
			} else if strArr[right] >= 'a' && strArr[right] <= 'z' {
				rightSub = strArr[right]
				break
			} else {
				right--
				if left >= right {
					return true
				}
			}
		}
		if left <= right && leftSub != rightSub {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		",.",
		"0P",
	}
	for _, testCase := range testCases {
		fmt.Println(testCase, isPalindrome(testCase))
	}
}
