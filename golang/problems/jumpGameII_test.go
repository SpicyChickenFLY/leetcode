package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type jumpGameIITestCase struct {
    nums []int
    expect int
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestJumpGameII(t *testing.T) {
    testCases := []jumpGameIITestCase {
        {
            nums:  []int{2, 3, 1, 1, 4} ,
            expect: 2,
        },
        {
            nums:  []int{2, 3, 0, 1, 4} ,
            expect: 2,
        },
        {
            nums:  []int{2, 0, 2, 4, 6, 0, 0, 3} ,
            expect: 3,
        },
        {
            nums:  []int{5,9,3,2,1,0,2,3,3,1,0,0} ,
            expect: 3,
        },
    }

    for index, testCase := range testCases {
        result := jump(testCase.nums)
        assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index + 1))
    }
}

