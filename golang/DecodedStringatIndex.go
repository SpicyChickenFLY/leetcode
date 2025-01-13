package main

import "fmt"
//import "strings"

func decodeAtIndex(S string, K int) string {
	s := []rune(S)
	k := 0
	pointer := 0
	for true {
		if s[pointer] > '0' && s[pointer] < '9' {
			print("number k:", k, " p:", pointer, "\n")
			for count := int(s[pointer]) - 1; count > 0; count-- {
				if k + pointer < K {
					print("repeat k:", k, " p:", pointer, "\n")
					k += pointer
				} else {
					print("return k:", k, " p:", pointer, "\n")
					return string(s[K - k + 1])
				}
			}
			pointer++
		} else {
			print("letter k:", k, " p:", pointer, "\n")
			if k == K {
				print("return k:", k, " p:", pointer, "\n")
				return string(s[pointer])
			}
			pointer++
			k++
		}
	}
	return "false"
}

func main(){
	var S = []string {
		"leet2code3",
		"ha22",
		"a2345678999999999999999",
	}

	for K := 0; K < 11; K++ {
		fmt.Println(S[0], K, decodeAtIndex(S[0], K), "\n")
	}
}