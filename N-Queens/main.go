package main

import "fmt"

func solveNQueens(n int) [][]string {
	// create board
	var board []string
	rowInitStr := ""
	for i := 0; i < n; i++ {
		rowInitStr += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, rowInitStr)
	}

	return recursion(board, 0)
}

func recursion(board []string, depth int) [][]string {
	result := [][]string{}
	if depth == len(board) {
		copied := make([]string, len(board))
		copy(copied, board)
		return append(result, copied)
	}

	for j := 0; j < len(board); j++ {
		if isValid(board, depth, j) {
			board[depth] = board[depth][:j] + "Q" + board[depth][j+1:]
			result = append(result, recursion(board, depth+1)...)
			board[depth] = board[depth][:j] + "." + board[depth][j+1:]
		}
	}

	return result
}

func isValid(board []string, x, y int) bool {
	// fmt.Print("isValid:", board, x, y)
	for i := 0; i < x; i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == 'Q' &&
				(j == y || x-i == y-j || x-i == j-y) {
				// fmt.Println(" false", x, y, i, j)
				return false
			}
		}
	}
	// fmt.Println(" true")
	return true
}

func main() {
	for testCase := 1; testCase <= 9; testCase++ {
		fmt.Println(testCase, solveNQueens(testCase))
	}
	// solveNQueens(4)
}
