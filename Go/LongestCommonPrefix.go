package main

import "fmt"
import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) > 0 {
		var minLen = len(strs[0])
		var position int = 0
		for pos, str := range strs {
			if len(str) < minLen {
				minLen = len(str)
				position = pos
			}
		}
		var maxWord = strs[position]
		
		for prefixLength := minLen; prefixLength > 0; prefixLength-- {
			prefix := maxWord[: prefixLength]
			var flag bool = true
			for _, str := range strs {
				if strings.Index(str, prefix) != 0{
					flag = false
					break
				}
			}
			if flag == true {
				return prefix
			}
		} 
		return ""
	} else {
		return ""
	}
}

func main() {
	var inputSet = [][]string{
		{"dog","racecar","car"},
		{"a"},
		{"fly", "flow", "flight"},
		{"ca", "a"},
	}
	for pos, input := range inputSet {
		
		fmt.Printf("%2d: %s\n", pos, longestCommonPrefix(input))
	} 
}