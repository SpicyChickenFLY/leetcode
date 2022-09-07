package main

import "fmt"

func validate(board [][]byte, i, j int) int {
	var flag int = 1023 // 0001 1111 1111
	// fmt.Printf("i:%d\tj:%d\tflag:%09b\n", i, j, flag)
	for k := 0; k < 9; k++ {
		if board[i][k] != '.' {
			// fmt.Printf("%q\n", board[i][k])
			flag &^= 1 << (board[i][k] - '1')
		}
		if board[k][j] != '.' {
			// fmt.Printf("%q\n", board[k][j])
			flag &^= 1 << (board[k][j] - '1')
		}
		row := i/3*3 + k/3
		col := j/3*3 + k%3
		if board[row][col] != '.' {
			// fmt.Printf("%q\n", board[row][col])
			flag &^= 1 << (board[row][col] - '1')
		}
	}
	// fmt.Printf("i:%d\tj:%d\tflag:%09b\n", i, j, flag)
	return flag
}

func dfs(board [][]byte, d int) bool {
	if d == 81 { // 9 * 9 all success
		return true
	}

	i := d / 9
	j := d % 9

	// skip fiiled number
	if board[i][j] != '.' {
		return dfs(board, d+1)
	}

	// test possible value
	flag := validate(board, i, j)
	for c := 0; c < 9; c++ {
		if flag&(1<<c) > 0 {
			fmt.Printf("i:%d\tj:%d\tflag:%09b\tc:%09b\tand:%09b\tval:%q\n", i, j, flag, 1<<c, flag&(1<<c), '1'+byte(c))
			board[i][j] = '1' + byte(c)
			if dfs(board, d+1) {
				return true
			}
		}
	}

	// wrong path
	fmt.Printf("error: i:%d\tj:%d\n", i, j)
	board[i][j] = '.'

	return false
}

func solveSudoku(board [][]byte) {
	dfs(board, 0)
}

func showBoard(board [][]byte) {
	fmt.Println("===================")
	fmt.Println("# 0 1 2 3 4 5 6 7 8")
	for lineID, line := range board {
		fmt.Print(lineID, " ")
		for _, element := range line {
			fmt.Printf("%c ", element)
		}
		fmt.Println()
	}
	fmt.Println("===================")
}

func main() {
	testCase := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	showBoard(testCase)
	// validate(testCase, 0, 6)
	solveSudoku(testCase)
	showBoard(testCase)
}
