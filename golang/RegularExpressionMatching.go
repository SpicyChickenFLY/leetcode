package main

import "fmt"
import "strings"

func isMatch_de(s string, p string) bool {
	sArray := []byte(s)
	sArray = append(sArray, '+')
	pArray := []byte(p)
	//normalize regular
	for pos := 0; pos < len(pArray) - 1; pos++ {
		if pArray[pos] == '*' && strings.Index(s, string(pArray[pos - 1])) == -1 {
			pArray = append(pArray[:pos - 1], pArray[pos + 1:]...)
		}
	}
	for pos := 0; pos < len(pArray) - 1; pos++ {
		if pArray[pos] == '*' && pArray[pos - 1] == pArray[pos + 1] {
			pArray[pos] = pArray[pos + 1]
			pArray[pos + 1] = '*'
		}else if pArray[pos] == '*' && pArray[pos + 1] == '*'{
			pArray = append(pArray[:pos], pArray[pos + 1:]...)
		}
	}
	fmt.Printf("s:%s, p:%s\n", sArray, pArray)
	var si  int = 0
	for pos := 0; pos < len(pArray); pos++ {
		fmt.Printf("Epoch-%d: %d-%s\n", pos, si, string(sArray[si]))
		if pos < len(pArray) - 1 && pArray[pos + 1] == '*' {
			for sArray[si] == pArray[pos] || (sArray[si] != '+' && pArray[pos] == '.') {
				si++
			}
			pos++ 
		}else if sArray[si] == pArray[pos] || (sArray[si] != '+' && pArray[pos] == '.') {
			si++
		}else {
			fmt.Printf("Error-%d: %d-%s\n", pos, si, string(sArray[si]))
			return false
		}
		fmt.Printf("Epoch end %d: %d-%s\n", pos, si, string(sArray[si]))
	}
	fmt.Printf("\nremain: %s\n", string(sArray[si]))
	return sArray[si] == '+'
}

func isMatch(s string, p string) bool {
	sArray := []byte(s)
	pArray := []byte(p)
	return true
}

func main() {
	var inputSet = [][]string{
		/*{"aa", "a"},
		{"aa", "a*"},
		{"aaa", "a*a"},
		{"a", "a*a"},
		{"ab", ".*"},
		{"aab", "c*a*b"},
		{"mississippi", "mis*is*p*."},
		{"ab", ".*c"},
		*/{"caaa", "cab*a*c*a"},
	}
	for pos, input := range inputSet {
		fmt.Printf("%d. s:%s, p:%s, %t\n\n", pos, input[0], input[1], isMatch(input[0], input[1]))
	}

}