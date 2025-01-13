package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximalRectangle(matrix [][]byte) int {
	height, width := len(matrix), len(matrix[0])
	// 记录坐标向右的最大长度
	flag := make([][]int, height)
	for i := 0; i < height; i++ {
		flag[i] = make([]int, width)
		if matrix[i][width - 1] == '1' {
			flag[i][width - 1] = 1
		}
		for j := width - 2; j >= 0; j-- {
			if matrix[i][j] == '1' {
				flag[i][j] = flag[i][j+1] + 1
			}
		}
	}

	// 计算每个坐标的最大的矩阵面积
	maxVal := flag[0][0]
	for j := 0; j < width; j++ {
		for i := 0; i < height; i++ {
			if (height - i) * (width - j) < maxVal {
				break // 长宽都太短了,接下来都不可能更大了
			}
			minVal := flag[i][j]
			maxVal = max(maxVal, minVal)
			for ii := i + 1; ii < height; ii++ {
				if flag[ii][j] == 0 {
					break // 高度中断
				}
				minVal = min(minVal, flag[ii][j])
				maxVal = max(maxVal, minVal * (ii - i + 1))
			}
		}
	}

	return maxVal
}

type maximalRectangleTestCase struct {
	matrix [][]byte
	expect  int
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMaximalRectangle(t *testing.T) {
	testCases := []maximalRectangleTestCase{
		{
			matrix: [][]byte{{'0', '1'}},
			expect:  1,
		},
		{
			matrix: [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}},
			expect:  6,
		},
		{
			matrix: [][]byte{{'0'}},
			expect:  0,
		},
		{
			matrix: [][]byte{{'1'}},
			expect:  1,
		},
	}

	for index, testCase := range testCases {
		result := maximalRectangle(testCase.matrix)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}
