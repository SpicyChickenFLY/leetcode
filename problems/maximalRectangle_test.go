package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximalRectangle(matrix [][]byte) int {
	// 记录坐标向右的最大长度
	// 这部分大家的算法基本差不多
	flag := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		flag[i] = make([]int, len(matrix[0]))
		if matrix[i][len(matrix[0]) - 1] == '1' {
			flag[i][len(matrix[0]) - 1] = 1
		}
		for j := len(matrix[0]) - 2; j >= 0; j-- {
			if matrix[i][j] == '1' {
				flag[i][j] = flag[i][j+1] + 1
			}
		}
	}

	// 计算每个坐标的最大的矩阵面积
	// 这部分好像有些算法用栈解决, 尝试优化一下
	maxVal := flag[0][0]
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			minVal := flag[i][j]
			if minVal > maxVal {
				maxVal = minVal
			}
			for ii := i + 1; ii < len(matrix); ii++ {
				if flag[ii][j] == 0 {
					break
				}
				if flag[ii][j] <= minVal {
					minVal = flag[ii][j]
				}
				if minVal * (ii - i + 1) > maxVal {
					maxVal = minVal * (ii - i + 1)
				}
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
