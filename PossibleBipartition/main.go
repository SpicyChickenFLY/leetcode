package main

import "fmt"

func dfs(groupMap map[int]bool, edges map[int][]int, node int, group bool) bool {
	if val, ok := groupMap[node]; ok {
		return val == group
	}

	groupMap[node] = group
	for _, neigh := range edges[node] {
		if !dfs(groupMap, edges, neigh, !group) {
			return false
		}
	}

	return true
}

func possibleBipartition(N int, dislikes [][]int) bool {
	groupMap := make(map[int]bool)
	edges := make(map[int][]int)

	for _, dislike := range dislikes {
		edges[dislike[0]] = append(edges[dislike[0]], dislike[1])
		edges[dislike[1]] = append(edges[dislike[1]], dislike[0])
	}

	for node := 1; node <= N; node++ {
		if _, ok := groupMap[node]; !ok {
			if !dfs(groupMap, edges, node, false) {
				return false
			}
		}
	}

	return true
}

func main() {
	testCases := [][][]int{
		{{4}, {1, 2}, {1, 3}, {2, 4}},
		{{3}, {1, 2}, {1, 3}, {2, 3}},
		{{5}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
		{{10},
			{4, 7}, {4, 8}, {5, 6}, {1, 6}, {3, 7}, {2, 5}, {5, 8}, {1, 2}, {4, 9}, {6, 1},
			{8, 1}, {3, 6}, {2, 1}, {9, 1}, {3, 9}, {2, 3}, {1, 9}, {4, 6}, {5, 7}, {3, 8},
			{1, 8}, {1, 7}, {2, 4}},
	}

	for _, testCase := range testCases {
		N := testCase[0][0]
		dislikes := testCase[1:]
		fmt.Println(testCase, possibleBipartition(N, dislikes))
	}
}
