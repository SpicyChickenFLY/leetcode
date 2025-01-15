package main

import "fmt"

func setZeroes(matrix [][]int) {
	var firstRowZero, firstColZero bool

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				continue
			}
			matrix[i][0] = 0
			matrix[0][j] = 0
			if i == 0 {
				firstRowZero = true
			}
			if j == 0 {
				firstColZero = true
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] != 0 {
			continue
		}
		for j := 1; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] != 0 {
			continue
		}
		for i := 1; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}

	if firstRowZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if firstColZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	testCases := [][][]int{
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		},
	}
	for _, testCase := range testCases {
		fmt.Println("test:", testCase)
		setZeroes(testCase)
		fmt.Println("result:", testCase)
	}
}
