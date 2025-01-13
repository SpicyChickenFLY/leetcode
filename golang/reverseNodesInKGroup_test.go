package golang

import (
	"testing"

    "fmt"

	// "github.com/stretchr/testify/assert"
)

type reverseNodeInKGroupTestCase struct {
    head *ListNode
    k int
    expect *ListNode
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestReverseNodeInKGroup(t *testing.T) {
    testCases := []reverseNodeInKGroupTestCase {
        {
            head: newList([]int{1, 2, 3, 4, 5}),
            k: 2,
            expect: newList([]int{2, 1, 4, 3, 5}),
        },
        {
            head: newList([]int{1, 2, 3, 4, 5}),
            k: 3,
            expect: newList([]int{3, 2, 1, 4, 5}),
        },
    }

    for _, testCase := range testCases {
        result := reverseKGroup(testCase.head, testCase.k)
        // fmt.Println(result)
        fmt.Println(strList(result))
        // assert.True(t, compareTwoList(testCase.expect, result), fmt.Sprintf("第%d个案例执行结果不一致", index + 1))
    }
}

