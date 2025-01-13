package main

import "fmt"

func generateMatrix(n int) [][]int {
	var res [][]int = make([][]int, n)
	for x := 0; x < n; x++ {
		res[x] = make([]int, n)
	}

	count := 1
	xStart, xEnd := 0, n-1
	yStart, yEnd := 0, n-1
	i, j := 0, 0
	for count <= n*n {
		for j = yStart; j <= yEnd; j++ {
			res[xStart][j] = count
			count++
		}
		xStart++
		for i = xStart; i <= xEnd; i++ {
			res[i][yEnd] = count
			count++
		}
		yEnd--
		for j = yEnd; j >= yStart; j-- {
			res[xEnd][j] = count
			count++
		}
		xEnd--
		for i = xEnd; i >= xStart; i-- {
			res[i][yStart] = count
			count++
		}
		yStart++
	}
	fmt.Println(i, j)
	return res
}

func main() {
	for i := 2; i < 5; i++ {
		res := generateMatrix(i)
		fmt.Println(res)
	}

}
