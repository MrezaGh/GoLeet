package main

import "fmt"

type WordDictionary struct {
	children map[rune]*WordDictionary
	valid    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		children: make(map[rune]*WordDictionary),
		valid:    false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	head := this
	for _, char := range word {
		if _, exist := head.children[char]; !exist {
			head.children[char] = &WordDictionary{
				children: make(map[rune]*WordDictionary),
				valid:    false,
			}
		}
		head = head.children[char]
	}
	head.valid = true
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.valid
	}

	if word[0] == '.' {
		for _, child := range this.children {
			if child.Search(word[1:]) {
				return true
			}
		}
		return false
	} else if node, exist := this.children[rune(word[0])]; exist {
		return node.Search(word[1:])
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad"), wordDictionary.Search("bad"),
		wordDictionary.Search("b.."), wordDictionary.Search(".ad"))
}
