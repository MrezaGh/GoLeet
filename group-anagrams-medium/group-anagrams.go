package main

import (
	"fmt"
	"slices"
)

type Trie struct {
	value  rune
	next   map[rune]*Trie
	active bool
	index  int
}

func (trie *Trie) InsertWordAndIndex(word []rune, index int) int {
	head := trie
	for _, c := range word {
		if next, ok := head.next[c]; ok {
			head = next
		} else {
			head.next[c] = &Trie{value: c, active: false, index: -1, next: make(map[rune]*Trie)}
			head = head.next[c]
		}
	}
	if head.active {
		return head.index
	} else {
		head.active = true
		head.index = index
		return index
	}
}

func groupAnagrams(strs []string) [][]string {
	trie := &Trie{value: 'h', next: make(map[rune]*Trie), active: false, index: -1}
	groupAna := make([][]string, 0)
	currentGroupIndex := 0
	for _, str := range strs {
		sortedStr := []rune(str)
		slices.Sort(sortedStr)
		groupIndex := trie.InsertWordAndIndex(sortedStr, currentGroupIndex)
		fmt.Printf("word: %s, currentgroupIndex: %d, found groupIndex: %d\n", str, currentGroupIndex, groupIndex)
		if groupIndex == currentGroupIndex {
			newGroup := []string{str}
			groupAna = append(groupAna, newGroup)
			currentGroupIndex += 1
			fmt.Printf("new anagram  group: %s\n", str)
		} else {
			groupAna[groupIndex] = append(groupAna[groupIndex], str)
			fmt.Printf("found anagram %s, in %v group\n", str, groupAna[groupIndex])
		}

	}

	return groupAna
}

func main() {
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs := []string{""}
	fmt.Println(groupAnagrams(strs))
}
