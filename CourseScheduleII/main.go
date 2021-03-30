package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := make([]int, numCourses)
	index := 0
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
				result[index] = i
				index++
			}
			if inDegree[i] > 0 {
				emptyFlag = false
			}
		}
		if !zeroFlag {
			if emptyFlag {
				return result
			}
			return []int{}
		}
	}
	return []int{}
}
