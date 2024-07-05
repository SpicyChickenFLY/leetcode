package problems

import (
	"testing"

    "fmt"

	"github.com/stretchr/testify/assert"
)

type reverseBetweenTestCase struct {
    head *ListNode
    left int
    right int
    expect *ListNode
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestReverseBetween(t *testing.T) {
    testCases := []reverseBetweenTestCase {
        {
            head: newList([]int{1, 2, 3}),
            left: 1,
            right: 3,
            expect: newList([]int{3, 2, 1}),
        },
        {
            head: newList([]int{1, 2, 3, 4, 5}),
            left: 2,
            right: 4,
            expect: newList([]int{1, 4, 3, 2, 5}),
        },
        {
            head: newList([]int{3, 5}),
            left: 1,
            right: 2,
            expect: newList([]int{5, 3}),
        },
        {
            head: newList([]int{5}),
            left: 1,
            right: 1,
            expect: newList([]int{5}),
        },
    }

    for index, testCase := range testCases {
        result := reverseBetween(testCase.head, testCase.left, testCase.right)
        assert.True(t, compareTwoList(testCase.expect, result), fmt.Sprintf("第%d个案例执行结果不一致", index + 1))
    }
}

