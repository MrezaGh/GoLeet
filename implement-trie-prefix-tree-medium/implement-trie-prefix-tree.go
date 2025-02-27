package main

import (
	"fmt"
)

type Trie struct {
	next  map[rune]*Trie
	valid bool
}

func Constructor() Trie {
	return Trie{
		next:  make(map[rune]*Trie),
		valid: false,
	}

}

func (this *Trie) Insert(word string) {
	head := this
	for _, char := range word {
		if _, exist := head.next[char]; !exist {
			head.next[char] = &Trie{
				next:  make(map[rune]*Trie),
				valid: false,
			}
		}
		head = head.next[char]

	}
	head.valid = true
}

func (this *Trie) Search(word string) bool {
	head := this
	for _, char := range word {
		if _, exist := head.next[char]; !exist {
			return false
		}
		head = head.next[char]
	}
	
	return head.valid
}

func (this *Trie) StartsWith(prefix string) bool {
	head := this
	for _, char := range prefix {
		if _, exist := head.next[char]; !exist {
			return false
		}
		head = head.next[char]
	}

	return true
}

func main() {

	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"), trie.Search("app"), trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
