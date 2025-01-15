package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
    customers []int
    grumpy []int
    minutes int
    expect int
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestIt(t *testing.T) {
    testCases := []TestCase {
        {
            customers:  []int{1, 0, 1, 2, 1, 1, 7, 5} ,
            grumpy: []int{0, 1, 0, 1, 0, 1, 0, 1},
            minutes: 3,
            expect: 16,
        },
        {
            customers:  []int{1} ,
            grumpy: []int{0},
            minutes: 1,
            expect: 1,
        },
        {
            customers:  []int{9, 10, 4, 5} ,
            grumpy: []int{1, 0, 1, 1},
            minutes: 1,
            expect: 19,
        },
    }

    for _, testCase := range testCases {
        result := maxSatisfied(testCase.customers, testCase.grumpy, testCase.minutes)
        assert.Equal(t, testCase.expect, result, "结果不一致")
    }
}

