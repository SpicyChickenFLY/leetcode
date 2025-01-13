package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sum := (C-A)*(D-B) + (G-E)*(H-F)
	left, right, up, down := max(A, E), min(G, C), min(D, H), max(F, B)
	if right > left && up > down {
		sum -= (right - left) * (up - down)
	}
	return sum
}
