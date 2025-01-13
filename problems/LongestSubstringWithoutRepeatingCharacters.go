package main

func lengthOfLongestSubstring(s string) int {
	record := make(map[byte]int)
	start := 0
	maxLen := 0

	for i, b := range []byte(s) {
		v, ok := record[b]
		if ok && v >= start {
			start = record[b] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		record[b] = i
	}
	return maxLen
}
