package main

const (
	living = -1
	dead   = 0
	live   = 1
	dying  = 2
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkNeighbour(board [][]int, x, y int) int {
	var result int
	xMin := max(x-1, 0)
	xMax := min(x+1, len(board)-1)
	yMin := max(y-1, 0)
	yMax := min(y+1, len(board[0])-1)

	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if i == x && j == y {
				continue
			}
			if board[i][j] > dead {
				result++
			}
		}
	}
	return result
}

func gameOfLife(board [][]int) {
	// first scan
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			neighbour := checkNeighbour(board, i, j)
			if board[i][j] == dead {
				if neighbour == 3 {
					board[i][j] = living
				}
			} else {
				if neighbour < 2 || neighbour > 3 {
					board[i][j] = dying
				}
			}
		}
	}
	// second scan
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == dying {
				board[i][j] = dead
			}
			if board[i][j] == living {
				board[i][j] = live
			}
		}
	}
}
