package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	result := make([]int, m*n)

	count := 0
	xStart, xEnd := 0, m-1
	yStart, yEnd := 0, n-1
	i, j := 0, 0
	for true {
		for j = yStart; j <= yEnd; j++ {
			result[count] = matrix[xStart][j]
			count++
			if count == m*n {
				return result
			}
		}
		xStart++
		for i = xStart; i <= xEnd; i++ {
			result[count] = matrix[i][yEnd]
			count++
			if count == m*n {
				return result
			}
		}
		yEnd--
		for j = yEnd; j >= yStart; j-- {
			result[count] = matrix[xEnd][j]
			count++
			if count == m*n {
				return result
			}
		}
		xEnd--
		for i = xEnd; i >= xStart; i-- {
			result[count] = matrix[i][yStart]
			count++
			if count == m*n {
				return result
			}
		}
		yStart++
	}
	return result
}

func main() {
	testCases := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
	}
	for _, testCase := range testCases {
		res := spiralOrder(testCase)
		fmt.Println(res)
	}
}
