package main

import "fmt"

func dfs(s string, dict map[string]bool, mem map[int][]string) []string {
	if result, ok := mem[len(s)]; ok {
		return result
	}
	result := []string{}
	for i := 1; i <= len(s); i++ {
		if dict[s[:i]] {
			// （终止条件）如果已经是最后一个子串了，那么将其本身也加入结果集
			if i == len(s) {
				result = append(result, s)
			} else {
				// 获取后续字符串的子结果
				subResults := dfs(s[i:], dict, mem)
				for _, subResult := range subResults {
					result = append(result, s[:i]+" "+subResult)
				}
			}

		}
	}
	mem[len(s)] = result
	return result
}

// 还是抄的同一个人的
func wordBreak(s string, wordDict []string) []string {
	// mem 用于记录某位点前子串是否正确分词
	mem := make(map[int][]string)

	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	return dfs(s, dict, mem)
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
