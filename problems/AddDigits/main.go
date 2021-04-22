package main

func addDigits(num int) int {
	for num/10 > 0 {
		temp := num % 10
		for num/10 > 0 {
			num /= 10
			temp += num % 10
		}
		num = temp
	}
	return num
}
