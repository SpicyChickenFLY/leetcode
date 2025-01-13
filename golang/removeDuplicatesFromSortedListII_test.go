package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type deleteDuplicatesTestCase struct {
    head *ListNode
    expect *ListNode
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestDeleteDuplicates(t *testing.T) {
    testCases := []deleteDuplicatesTestCase {
        {
            head: newList([]int{1, 2, 3, 3, 4, 4, 5}),
            expect: newList([]int{1, 2, 5}),
        },
        {
            head: newList([]int{1, 1, 1, 2, 3}),
            expect: newList([]int{2, 3}),
        },
        {
            head: newList([]int{1, 1, 1}),
            expect: newList([]int{}),
        },
        {
            head: newList([]int{}),
            expect: newList([]int{}),
        },
    }

    for _, testCase := range testCases {
        result := deleteDuplicates(testCase.head)
        assert.True(t, compareTwoList(testCase.expect, result), "结果不一致")
    }
}

