package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strArr := strings.Split(s, " ")
	result := []string{}
	for i := len(strArr) - 1; i >= 0; i-- {
		if len(strArr[i]) != 0 {
			result = append(result, strArr[i])
		}
	}
	return strings.Join(result, " ")
}

func main() {
	testCases := []string{
		"the sky is blue",
		"   hello world  ",
		"a good   example",
	}
	for _, testCase := range testCases {
		fmt.Println(testCase)
		fmt.Println(reverseWords(testCase))
	}

}
