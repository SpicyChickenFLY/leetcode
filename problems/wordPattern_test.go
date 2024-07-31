package problems

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func wordPattern(pattern string, s string) bool {
	patternWordMap := map[byte]string{}
	wordPatternMap := map[string]byte{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if recordWord, exists := patternWordMap[pattern[i]]; !exists {
			patternWordMap[pattern[i]] = words[i]
		} else if recordWord != words[i] {
			return false
		}
		if recordPat, exists := wordPatternMap[words[i]]; !exists {
			wordPatternMap[words[i]] = pattern[i]
		} else if recordPat != pattern[i] {
			return false
		}
	}
	return true
}

type wordPatternTestCase struct {
	pattern string
	s       string
	expect  bool
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestWordPattern(t *testing.T) {
	testCases := []wordPatternTestCase{
		{
			pattern: "jquery",
			s:       "jquery",
			expect:  false,
		},
		{
			pattern: "abba",
			s:       "dog cat cat dog",
			expect:  true,
		},
		{
			pattern: "abba",
			s:       "dog cat cat fish",
			expect:  false,
		},
		{
			pattern: "abba",
			s:       "dog dog dog dog",
			expect:  false,
		},
		{
			pattern: "aaa",
			s:       "dog dog dog dog",
			expect:  false,
		},
	}

	for index, testCase := range testCases {
		result := wordPattern(testCase.pattern, testCase.s)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}
