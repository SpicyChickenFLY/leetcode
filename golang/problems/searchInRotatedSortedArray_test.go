package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SearchInRotatedSortedArrayTestCase struct {
    nums []int
    target int
    expect int
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSearchInRotatedSortedArray(t *testing.T) {
    testCases := []SearchInRotatedSortedArrayTestCase {
        {
            nums:  []int{3, 5, 1},
            target: 3,
            expect: 0,
        },
        {
            nums:  []int{1, 3},
            target: 3,
            expect: 1,
        },
        {
            nums:  []int{1, 3, 5},
            target: 1,
            expect: 0,
        },
        {
            nums:  []int{1},
            target: 0,
            expect: -1,
        },
        {
            nums:  []int{4, 5, 6, 7, 0, 1, 2},
            target: 0,
            expect: 4,
        },
        {
            nums:  []int{4, 5, 6, 7, 0, 1, 2} ,
            target: 1,
            expect: 5,
        },
    }

    for index, testCase := range testCases {
        result := search(testCase.nums, testCase.target)
        assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个测试案例结果不一致", index + 1))
    }
}

