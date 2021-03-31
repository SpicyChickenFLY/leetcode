package main

import "fmt"

type wordNode struct {
	Next       map[rune]*wordNode
	Terminated bool
}

func newNode() *wordNode {
	return &wordNode{
		Next:       make(map[rune]*wordNode),
		Terminated: false,
	}
}

type WordDictionary struct {
	dictTree *wordNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{dictTree: newNode()}
}

func (this *WordDictionary) AddWord(word string) {
	temp := this.dictTree
	for _, letter := range word {
		if subNode, ok := temp.Next[letter]; ok {
			temp = subNode
		} else {
			subNode = newNode()
			temp.Next[letter] = subNode
			temp = subNode
		}
	}
	temp.Terminated = true
}

func dfs(word string, dictTree *wordNode) bool {
	temp := dictTree
	if len(word) == 1 {
		for _, node := range temp.Next {
			if node.Terminated {
				return true
			}
		}
		return false
	}
	for _, node := range temp.Next {
		if search(word[1:], node) {
			return true
		}
	}
	return false
}

func search(word string, dictTree *wordNode) bool {
	temp := dictTree
	for i, letter := range word {
		if letter == '.' {
			return dfs(word[i:], temp)
		} else if subNode, ok := temp.Next[letter]; ok {
			temp = subNode
			continue
		}
		return false
	}
	return temp.Terminated
}

func (this *WordDictionary) Search(word string) bool {
	return search(word, this.dictTree)
}

func main() {
	wordDict := Constructor()
	wordDict.AddWord("at")
	wordDict.AddWord("and")
	wordDict.AddWord("an")
	wordDict.AddWord("add")
	fmt.Println(wordDict.Search("a"))
	fmt.Println(wordDict.Search(".at"))
	wordDict.AddWord("bat")
	fmt.Println(wordDict.Search(".at"))
	fmt.Println(wordDict.Search("an."))
	fmt.Println(wordDict.Search("a.d."))
	fmt.Println(wordDict.Search("b."))
	fmt.Println(wordDict.Search("a.d"))
	fmt.Println(wordDict.Search("."))

	// wordDict.AddWord("bad")
	// wordDict.AddWord("dad")
	// wordDict.AddWord("mad")
	// fmt.Println(wordDict.Search("pad"))
	// fmt.Println(wordDict.Search("bad"))
	// fmt.Println(wordDict.Search(".ad"))
	// fmt.Println(wordDict.Search("b.."))

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
