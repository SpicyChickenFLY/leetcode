package main

func trailingZeroes(n int) int {
	result := 0
	// 除一次不够，像25的数字中存在两个5
	for n >= 5 {
		n /= 5
		result += n
	}
	return result
}

func main() {
}
