package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Range struct {
	leftBorder    int // 记录划定区域的左边界
	minHeight int // 划定区域的最小高度(左边界的高度, 因为该区域中高度是单调递增)
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]Range, 0)

    // 从左向右遍历各个高度
	for rightDivider, height := range heights {
		start := rightDivider
        // 如果栈不空,且当前高度比栈顶高度小,则划定了一个局部区域
		for len(stack) != 0 && height < stack[len(stack)-1].minHeight {
            // 区域的横跨面积为 <最小高度> * <区域宽度>
			maxArea = max(maxArea, stack[len(stack)-1].minHeight*(rightDivider-stack[len(stack)-1].leftBorder))
            // 记录局部区域的左边界位置
			start = stack[len(stack)-1].leftBorder
            // 抛出栈元素,继续判断下一个栈顶元素是否划定局部区域
			stack = stack[:len(stack)-1]
		}
        // 不能继续划分局部区域,将当前高度作为区域最小高度,并且变为区域的左边界
		stack = append(stack, Range{start, height})
	}

    // 遍历完成,可以看作是当前高度为0,则再次划定局部区域
	for _, bar := range stack {
		maxArea = max(maxArea, bar.minHeight*(len(heights)-bar.leftBorder))
	}

	return maxArea
}

func TestLargestRectangleInHistogram(t *testing.T) {
	testCases := []largestRectangleInHistogramTestCase{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			expect:  10,
		},
		{
			heights: []int{2, 4},
			expect:  4,
		},
	}

	for index, testCase := range testCases {
		result := largestRectangleArea(testCase.heights)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}

type largestRectangleInHistogramTestCase struct {
	heights []int
	expect  int
}
