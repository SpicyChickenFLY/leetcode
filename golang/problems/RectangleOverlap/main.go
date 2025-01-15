package main

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

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	left, right, up, down :=
		max(rec1[0], rec2[0]),
		min(rec1[2], rec2[2]),
		min(rec1[3], rec2[3]),
		max(rec1[1], rec2[1])
	return right > left && up > down
}
