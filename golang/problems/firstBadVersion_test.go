package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type firstBadVersionTestCase struct {
    n int
    bad int
    expect int
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFirstBadVersion(t *testing.T) {
    testCases := []firstBadVersionTestCase {
        {
            n: 5,
            bad: 4,
            expect: 4,
        },
    }

    for index, testCase := range testCases {
        badVersion = testCase.bad
        result := firstBadVersion(testCase.n)
        assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index + 1))
    }
}

