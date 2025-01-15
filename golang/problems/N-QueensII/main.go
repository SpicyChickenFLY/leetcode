package main

import "fmt"

func totalNQueens(n int) int {
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

func recursion(board []string, depth int) int {
	result := 0
	if depth == len(board) {
		copied := make([]string, len(board))
		copy(copied, board)
		return 1
	}

	for j := 0; j < len(board); j++ {
		if isValid(board, depth, j) {
			board[depth] = board[depth][:j] + "Q" + board[depth][j+1:]
			result += recursion(board, depth+1)
			board[depth] = board[depth][:j] + "." + board[depth][j+1:]
		}
	}

	return result
}

func isValid(board []string, x, y int) bool {
	for i := 0; i < x; i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == 'Q' &&
				(j == y || x-i == y-j || x-i == j-y) {
				return false
			}
		}
	}
	return true
}

func main() {
	for testCase := 1; testCase <= 9; testCase++ {
		fmt.Println(testCase, totalNQueens(testCase))
	}
}
