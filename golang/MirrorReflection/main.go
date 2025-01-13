package main

func gcd(a, b int) int {
	if a < b { //交换两个数，使大数放在a上
		a, b = b, a
	}
	for b != 0 { //利用辗除法，直到b为0为止
		a, b = b, a%b
	}
	return a
}

func isOddAdd(num int) bool {
	return num%2 != 0
}

func mirrorReflection(p int, q int) int {
	x := gcd(p, q)
	if isOddAdd(q / x) {
		if isOddAdd(p / x) {
			return 1
		} else {
			return 2
		}
	} else {
		return 0
	}
}
