package main

import "fmt"

type UnionFind struct {
	m map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		m: make(map[int]int),
	}
}

func (uf *UnionFind) Find(x int) int {
	for uf.m[x] != 0 {
		x = uf.m[x]
	}
	return x
}

func (uf *UnionFind) Union(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x > y {
		x, y = y, x
	}
	if x != y {
		uf.m[y] = x
	}
}

func isLeftRightConnected(left, right int) bool {
	return (left == 1 || left == 4 || left == 6) &&
		(right == 1 || right == 3 || right == 5)
}

func isUpDownConnected(up, down int) bool {
	return (up == 2 || up == 3 || up == 4) &&
		(down == 2 || down == 5 || down == 6)
}

// like two-pass connected-component
func hasValidPath(grid [][]int) bool {
	h := len(grid)
	w := len(grid[0])
	colorMap := make([][]int, h)
	for i := range colorMap {
		colorMap[i] = make([]int, w)
	}
	colorCount := 1
	colorRelation := NewUnionFind()
	// traverse
	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			if x > 0 && isUpDownConnected(grid[x-1][y], grid[x][y]) {
				colorMap[x][y] = colorMap[x-1][y]
			}
			if y > 0 && isLeftRightConnected(grid[x][y-1], grid[x][y]) {
				if colorMap[x][y] != 0 {
					colorRelation.Union(colorMap[x][y], colorMap[x][y-1])
				} else {
					colorMap[x][y] = colorMap[x][y-1]
				}
			}
			if colorMap[x][y] == 0 {
				colorMap[x][y] = colorCount
				colorCount++
			}
			// fmt.Println(colorMap, colorCount, colorRelation.m)
		}
	}
	// check neighbour connected
	return colorRelation.Find(colorMap[0][0]) == colorRelation.Find(colorMap[h-1][w-1])
}

func main() {
	testCases := [][][]int{
		{{2, 4, 3}, {6, 5, 2}},              // true
		{{1, 2, 1}, {1, 2, 1}},              // false
		{{1, 1, 2}},                         // false
		{{1, 1, 1, 1, 1, 1, 3}},             // true
		{{2}, {2}, {2}, {2}, {2}, {2}, {6}}, // true
	}
	for _, testCase := range testCases {
		fmt.Println(testCase, hasValidPath(testCase))
	}
}
