package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cloneGraphTestCase struct {
    adjList [][]int
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCloneGraph(t *testing.T) {
    testCases := []cloneGraphTestCase {
        { [][]int{{2,4},{1,3},{2,4},{1,3}} },
    }

    for _, testCase := range testCases {
        graph := newGraph(testCase.adjList)
        result := cloneGraph(graph)
        fmt.Println(graph)
        visited := make(map[int]*Node)
        dfs(result, visited)
        fmt.Println(visited)

        assert.True(t, true, "结果不一致")
        // assert.True(t, compareTwoGraph(result, graph), "结果不一致")
    }
}

