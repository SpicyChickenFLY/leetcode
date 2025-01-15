package main

// 通过DFS的思想进行克隆
func dfsClone(graph *Node, visited map[*Node]*Node) *Node {
	// 克隆并记录本节点
	newNode := &Node{graph.Val, make([]*Node, len(graph.Neighbors))}
	visited[graph] = newNode
	// 遍历本节点所有邻节点
	for index, neighbor := range graph.Neighbors {
		if visitedNode, ok := visited[neighbor]; ok {
			// 记录过的节点直接记录信息
			newNode.Neighbors[index] = visitedNode
		} else {
			// 未记录过的节点继续递归克隆子节点后再记录信息
			newNode.Neighbors[index] = dfsClone(neighbor, visited)
		}
	}
	return newNode
}
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
    visited := make(map[*Node]*Node)
	return dfsClone(node, visited)
}

// BFS
func cloneGraphWithBFS(node *Node) *Node {
    if node == nil { return nil}
    visited := map[*Node]*Node{ node: { Val: node.Val } }
    queue := []*Node{node}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]

        for _, neighbor := range curr.Neighbors {
            if _, ok := visited[neighbor]; !ok {
                visited[neighbor] = &Node{ Val: neighbor.Val}
                queue = append(queue, neighbor)
            }
            visited[curr].Neighbors = append(visited[curr].Neighbors, visited[neighbor])
        }
    }
    return visited[node]
}
