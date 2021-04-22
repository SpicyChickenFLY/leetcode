package main

import "fmt"

// 抄作业的
func wordBreak(s string, wordDict []string) bool {
	// mem 用于记录某位点前子串是否正确分词
	mem := make([]bool, len(s)+1)
	mem[0] = true

	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if mem[j] && dict[s[j:i]] {
				mem[i] = true // 拼接子串并更新位点信息
				break
			}
		}
	}
	return mem[len(s)] // 判断最后一个字符位点是否正确
}

func main() {
	stringTestCases := []string{
		"bb",
		"leetcode",
		"applepenapple",
		"catsandog",
	}
	wordDictTestCases := [][]string{
		{"a", "b", "bbb", "bbbb"},
		{"leet", "code"},
		{"apple", "pen"},
		{"cats", "dog", "sand", "and", "cat"},
	}
	for i := 0; i < len(stringTestCases); i++ {
		fmt.Println(
			stringTestCases[i],
			wordDictTestCases[i],
			wordBreak(stringTestCases[i], wordDictTestCases[i]))
	}
}
