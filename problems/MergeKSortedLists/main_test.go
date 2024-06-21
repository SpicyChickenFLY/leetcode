package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
    lists []*ListNode
    expect *ListNode
}

func compareTwoList(l1, l2 *ListNode) bool {
    for l1 != nil || l2 != nil {
        if l1 == nil || l2 == nil {
            return false
        }
        if l1.Val != l2.Val {
            return false
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return true
}

func strList(head *ListNode) {
    for head != nil {
        fmt.Printf("->%d", head.Val)
        head = head.Next
    }
    fmt.Printf("\n")
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMerge(t *testing.T) {
    testCases := []TestCase {
        {
            lists:  newLists([][]int{ {1, 4, 5}, {1, 3, 4}, {2, 6} }) ,
            expect: newList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
        },
        {
            lists:  newLists([][]int{ {1}, {0} }) ,
            expect: newList([]int{0, 1}),
        },
        {
            lists : newLists([][]int{ {1, 2, 3}, {4, 5, 6, 7} }),
            expect: newList([]int{1, 2, 3, 4, 5, 6, 7}),
        },
        {
            lists : newLists([][]int{ {1, 4, 5}, {0, 2} }),
            expect: newList([]int{0, 1, 2, 4, 5}),
        },
    }

    for _, testCase := range testCases {
        result := mergeKLists(testCase.lists)
        strList(result)
        assert.True(t, compareTwoList(result, testCase.expect), "结果不一致")
    }
}

