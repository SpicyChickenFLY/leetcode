package problems

/**
 * Definition for a Node.
 */
type Node struct {
    Val int
    Neighbors []*Node
}

func newGraph(adjList [][]int) *Node {
    arr := []*Node{}
    for index := range adjList {
        arr = append(arr, &Node{index + 1, []*Node{}})
	}
	for index, node := range adjList {
        for _, adj := range node {
            arr[index].Neighbors = append(arr[index].Neighbors, arr[adj - 1])
        }
	}
    if len(arr) > 0 {
        return arr[0]
    } else {
        return nil
    }
}

func dfs(graph *Node, visited map[int]*Node) {
	if graph == nil {
		return
	}
	visited[graph.Val] = graph

	// for each of the adjacent vertices, call the function recursively
	// if it hasn't yet been visited
	for _, neighbor := range graph.Neighbors {
		if visited[neighbor.Val] != nil {
            continue
		}
        dfs(neighbor, visited)
	}
}

func graphToAdjList(graph *Node) string {
    visited := make(map[int]*Node)
    dfs(graph, visited)
    return ""
}
