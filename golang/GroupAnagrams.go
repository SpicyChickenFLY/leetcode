package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	dict := make(map[[26]uint8]int)

	for _, str := range strs {
		key := [26]uint8{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}

		if index, ok := dict[key]; ok {
			result[index] = append(result[index], str)
		} else {
			dict[key] = len(dict)
			result = append(result, []string{str})
		}
	}
	return result
}

func main() {
	fmt.Println("")
	testCase := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	groupAnagrams(testCase)
}
