package problems

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	longest := 0
	length := 0
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			continue
		} else if nums[i] == last+1 {
			length++
		} else {
			longest = max(longest, length+1)
			length = 0
		}
		last = nums[i]
	}
	longest = max(longest, length+1)
	return longest
}

type longestConsecutiveTestCase struct {
	nums   []int
	expect int
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestLongestConsecutive(t *testing.T) {
	testCases := []longestConsecutiveTestCase{
		{
			nums:   []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expect: 9,
		},
		{
			nums:   []int{100, 4, 200, 1, 3, 2},
			expect: 4,
		},
	}

	for index, testCase := range testCases {
		result := longestConsecutive(testCase.nums)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}
