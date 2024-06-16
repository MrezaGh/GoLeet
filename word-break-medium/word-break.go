package main

import (
	"errors"
	"fmt"
)

func main() {
	//s := "leetcode"
	//wordDict := []string{"leet", "code"}

	//s := "applepenapple"
	//wordDict := []string{"apple", "pen"}

	//s := "catsandog"
	//wordDict := []string{"cats", "dog", "sand", "and", "cat"}

	//s := "sandal"
	//wordDict := []string{"san", "al", "sand"}

	//s := "a"
	//wordDict := []string{"a"}

	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	//s := "catsandogcat"
	//wordDict := []string{"cats", "dog", "sand", "and", "cat", "an"}
	fmt.Println(revisedWordBreak(s, wordDict))
	//fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	wordsMap := make(map[string]struct{})
	for _, value := range wordDict {
		wordsMap[value] = struct{}{}
	}
	checked := make(map[string]int)
	_, err := wordBreakDP(s, wordsMap, checked)
	//result, err := wordBreakDP(s, wordsMap, checked)
	//fmt.Println(result)
	//fmt.Println(err)
	if err == nil {
		return true
	}
	return false
}

func wordBreakDP(s string, wordMap map[string]struct{}, checked map[string]int) ([]string, error) {
	if s == "" {
		return []string{}, nil
	}
	//fmt.Println("checking:", s, "checked:", checked)
	for i := 0; i < len(s); i++ {
		if _, ok := wordMap[s[:i+1]]; ok {
			//checked
			//if index, ok := checked[s[:i+1]]; ok && index == len(s)-i {
			//if index, ok := checked[s[i:]]; ok && index == i {
			if _, ok := checked[s[i:]]; ok {
				fmt.Println("skip", s[:i+1])
				//already checked that is did not work
				continue
			}
			segmentedParts, err := wordBreakDP(s[i+1:], wordMap, checked)
			if err != nil {
				fmt.Println(fmt.Sprintf("found <%s>; but did not work", s[:i+1]))
				checked[s[i:]] = i
				//checked[s[:i+1]] = i
				continue
				//return nil, errors.New(fmt.Sprintf("found %s; but did not work", s[:i]))
			}
			//fmt.Println("debug. found: ", segmentedParts)
			return append(segmentedParts, s[:i+1]), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("failed to find anything for <%s> in map:<%s>", s, wordMap))

}

func revisedWordBreak(s string, wordDict []string) bool {
	wordsMap := make(map[string]struct{})
	for _, value := range wordDict {
		wordsMap[value] = struct{}{}
	}
	possible := make([]bool, len(s))
	checkValidWords(0, s, possible, wordsMap)
	//fmt.Println(possible)
	for i := 0; i < len(s)-1; i++ {
		if possible[i] {
			checkValidWords(i+1, s, possible, wordsMap)
			//fmt.Println("index:", i+1, " possible:", possible)
		}
	}

	//println(possible)
	return possible[len(possible)-1]
}

func checkValidWords(index int, s string, possible []bool, wordsMap map[string]struct{}) {
	minLen := 21
	if len(s) < index+minLen {
		minLen = len(s) - index
	}

	for i := index; i < index+minLen; i++ {
		if _, ok := wordsMap[s[index:i+1]]; ok {
			//fmt.Println(s[index : i+1])
			possible[i] = true
		}
	}
}
