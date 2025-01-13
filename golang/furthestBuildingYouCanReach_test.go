package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type furthestBuildingYouCanReachTestCase struct {
    heights []int
    bricks int
    ladders int
    expect int
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFurthestBuildingYouCanReach(t *testing.T) {
    testCases := []furthestBuildingYouCanReachTestCase {
        {
            heights:  []int{1, 5, 8, 8} ,
            bricks: 3,
            ladders: 1,
            expect: 3,
        },
        {
            heights:  []int{14, 3, 19, 3} ,
            bricks: 17,
            ladders: 0,
            expect: 3,
        },
        {
            heights:  []int{4, 12, 2, 7, 3, 18, 20, 3, 19} ,
            bricks: 10,
            ladders: 2,
            expect: 7,
        },
        {
            heights:  []int{4, 2, 7, 6, 9, 14, 12} ,
            bricks: 5,
            ladders: 1,
            expect: 4,
        },
    }

    for index, testCase := range testCases {
        result := furthestBuilding(testCase.heights, testCase.bricks, testCase.ladders)
        assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index + 1))
    }
}

