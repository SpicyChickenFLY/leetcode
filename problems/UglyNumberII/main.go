package main

import (
	"fmt"
)

// 抄作业的，说实话还不是太懂
func nthUglyNumber(n int) int {
	arr := make([]int, n)
	arr[0] = 1
	pf2, pf3, pf5 := 0, 0, 0
	for i := 1; i < len(arr); i++ {
		t2, t3, t5 := 2*arr[pf2], 3*arr[pf3], 5*arr[pf5]
		v := min(t2, min(t3, t5))
		arr[i] = v
		// 当出现下一个不是ugly number时，会导致这里会有多个值自增
		if v == t2 {
			pf2++
		}
		if v == t3 {
			pf3++
		}
		if v == t5 {
			pf5++
		}
	}
	return arr[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(i, nthUglyNumber(i))
	}
}
