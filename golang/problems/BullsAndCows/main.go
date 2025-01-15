package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	counterBull := 0
	counterCow := 0
	counter := map[byte]int{
		'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0,
	}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			counterBull++
		} else {
			if counter[secret[i]] < 0 {
				counterCow++
			}
			if counter[guess[i]] > 0 {
				counterCow++
			}
			counter[secret[i]]++
			counter[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", counterBull, counterCow)
}

func main() {
	testCases := [][]string{
		{"1807", "7810"},
		{"1123", "0111"},
		{"1", "0"},
		{"1", "1"},
	}
	for _, testCase := range testCases {
		fmt.Println(testCase, getHint(testCase[0], testCase[1]))
	}
}
