package main

// func DFS

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	topoMap := make(map[int][]int)
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
		topoMap[pre[1]] = append(topoMap[pre[1]], pre[0])
	}
	for true {
		zeroFlag := false
		emptyFlag := true
		for i := 0; i < numCourses; i++ {
			if inDegree[i] == 0 {
				zeroFlag = true
				inDegree[i] = -1
				for _, outEdge := range topoMap[i] {
					inDegree[outEdge]--
				}
			}
			if inDegree[i] > 0 {
				emptyFlag = false
			}
		}
		if !zeroFlag {
			return emptyFlag
		}
	}
	return false
}

func main() {

}
