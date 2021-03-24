package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]byte, len(num1)+len(num2))
	for i := len(num1) + len(num2) - 1; i >= 0; i-- {
		result[i] = '0'
	}
	for i := len(num1) - 1; i >= 0; i-- {
		var carry byte = '0' - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			temp := (result[i+j+1] - '0') + (num2[j]-'0')*(num1[i]-'0') + carry
			carry = temp / 10
			result[i+j+1] = temp%10 + '0'
		}
		if carry > 0 {
			result[i] = carry + '0'
		}
	}
	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}

func main() {
	testCases := [][]string{
		{"2", "3", "6"},
		{"24", "9", "216"},
		{"498828660196", "840477629533", "419254329864656431168468"},
	}

	for _, testCase := range testCases {
		fmt.Println(testCase[0], testCase[1], testCase[2], multiply(testCase[0], testCase[1]))
	}
}
