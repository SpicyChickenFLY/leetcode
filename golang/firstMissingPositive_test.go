package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type firstMissingPositiveTestCase struct {
    nums []int
    expect int
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFirstMissingPositive(t *testing.T) {
    testCases := []firstMissingPositiveTestCase {
        {
            nums:  []int{3, 4, -1, 1} ,
            expect: 2,
        },
        {
            nums:  []int{1} ,
            expect: 2,
        },
        {
            nums:  []int{0, 2, 2, 1, 1} ,
            expect: 3,
        },
        {
            nums:  []int{1, 2, 0} ,
            expect: 3,
        },
        {
            nums:  []int{7, 8, 9, 11, 12} ,
            expect: 1,
        },
        {
            nums:  []int{1, 1} ,
            expect: 2,
        },
    }

    for index, testCase := range testCases {
        result := firstMissingPositive(testCase.nums)
        assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index + 1))
    }
}

