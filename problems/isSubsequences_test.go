package problems

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
func isSubsequence(s string, t string) bool {
	tLen, sLen := len(t), len(s)
	// 可以通过给矩阵多加一行和一列来避免多余的首行和首列判断,空间换时间
	matrix := make([][]int, sLen+1)
	// 初始化矩阵首行
	matrix[0] = make([]int, tLen+1)
	for j := 0; j <= tLen - sLen; j++ {
		matrix[0][j] = 1
	}

    // 目标字串逐字(也就是矩阵行)匹配
	for i := 1; i <= sLen; i++ {
		matrix[i] = make([]int, tLen + 1)
		for j := 1; j <= tLen; j++ {
            // 目标字串位置保证在源字串内部
			if j >= i && j <= tLen - sLen + i {
                if  s[i-1] == t[j-1] {
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
	return matrix[sLen][tLen] > 0
}

type isSubsequenceTestCase struct {
	s      string
	t      string
	expect bool
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestIsSubsequence(t *testing.T) {
	testCases := []isSubsequenceTestCase{
		{
			s:      "rabbit",
			t:      "rabbbit",
			expect: true,
		},
	}

	for index, testCase := range testCases {
		result := isSubsequence(testCase.s, testCase.t)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}
