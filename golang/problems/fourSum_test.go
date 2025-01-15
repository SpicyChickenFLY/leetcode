package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	arr := make([]int, 4)
	return kSum(4, 0, int64(target), nums, [][]int{}, arr)
}

// k should be less than 4
func kSum(k int, start int, target int64, nums []int, res [][]int, record []int) [][]int {
	if k > 2 {
		for i := start; i < len(nums)-k+1; i++ {
			if i > start && nums[i] == nums[i-1] {
				// 上一个数字已经探索过所有可能了, 相同数字没有必要再次处理
				continue
			}
			record[len(record)-k] = nums[i]
			// 假定这个数字是一种可能的倒数第k个值, 递归寻找其余的数字
			res = kSum(k-1, i+1, target-int64(nums[i]), nums, res, record)
		}
	} else {
		// 递归终结, 记录左右边界位置
		l, r := start, len(nums)-1
		// 循环直至左右边界重合
		for l < r {
			sum := int64(nums[l]) + int64(nums[r])
			if sum < target {
				l += 1 // 和值小于目标, 左边界右移
			} else if sum > target {
				r -= 1 // 和值大于目标, 右边界左移
			} else {
				// 和值匹配目标, 记录 可能性
				res = append(res, []int{record[0], record[1], nums[l], nums[r]}) // 如果需要满足k,需要改造
				// 这个数字已经探索过所有可能了, 相同数字没有必要再次处理
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}
	return res
}

type fourSumTestCase struct {
	nums      []int
	target      int
	expect [][]int
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFourSum(t *testing.T) {
	testCases := []fourSumTestCase{
		{
			nums:      []int{1,0,-1,0,-2,2},
			target:      0,
			expect: [][]int{{-2,-1,1,2}, {-2,0,0,2}, {-1,0,0,1}},
		},
	}

	for index, testCase := range testCases {
		result := fourSum(testCase.nums, testCase.target)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}
