package main

import (
	"fmt"
	"reflect"
)

func main() {
	//s := "anagram"
	s := "rat"
	//t := "nagaram"
	t := "car"
	fmt.Println(isAnagram(s, t))

}

func isAnagram(s string, t string) bool {
	sMap, tMap := make(map[string]int), make(map[string]int)
	for _, c := range s {
		sMap[string(c)] += 1
	}
	for _, c := range t {
		tMap[string(c)] += 1
	}

	return reflect.DeepEqual(sMap, tMap)
}
