package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//     b a b g b a g
//   1 1 1 1 1 0 0 0 // 不需要全设1是因为最后长度不够就没必要匹配了
// b 0 1 1 2 2 3 0 0 // 不匹配取左侧结果, 匹配取左侧加左上
// a 0 0 1 1 1 1 4 0
// g 0 0 0 0 1 1 1 5
func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	// 可以通过给矩阵多加一行和一列来避免多余的首行和首列判断,空间换时间
	matrix := make([][]int, tLen+1)
	// 初始化矩阵首行
	matrix[0] = make([]int, sLen+1)
	for j := 0; j <= sLen - tLen; j++ {
		matrix[0][j] = 1
	}

    // 目标字串逐字(也就是矩阵行)匹配
	for i := 1; i <= tLen; i++ {
		matrix[i] = make([]int, sLen + 1)
		for j := 1; j <= sLen; j++ {
            // 目标字串位置保证在源字串内部
			if j >= i && j <= sLen - tLen + i {
                if  t[i-1] == s[j-1] {
                    // 记录当前字已有可能结果和前面字匹配可能结果
                    matrix[i][j] = matrix[i][j-1] + matrix[i-1][j-1]
                } else {
                    // 不匹配, 仅记录当前字已有可能结果
                    matrix[i][j] = matrix[i][j-1]
                }
            }
		}
	}
    // 最终可能的匹配数就是矩阵最后一个字的匹配结果
	return matrix[tLen][sLen]
}

type numDistinctTestCase struct {
	s      string
	t      string
	expect int
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNumDistinct(t *testing.T) {
	testCases := []numDistinctTestCase{
		{
			s:      "rabbbit",
			t:      "rabbit",
			expect: 3,
		},
	}

	for index, testCase := range testCases {
		result := numDistinct(testCase.s, testCase.t)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}
