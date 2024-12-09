package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := buildGraph(beginWord, wordList)
	//fmt.Println(graph)
	return findPath(beginWord, endWord, graph)
}

type queueItem struct {
	node       string
	pathLength int
}

func findPath(start, end string, graph map[string][]string) int {
	visited := make(map[string]bool)
	queue := make([]queueItem, 0)
	queue = append(queue, queueItem{start, 1})
	for len(queue) > 0 {
		head, pathLength := queue[0].node, queue[0].pathLength
		visited[head] = true
		if head == end {
			//fmt.Println("found it")
			return pathLength
		}
		for _, neighbor := range graph[head] {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, queueItem{neighbor, pathLength + 1})
			}
		}
		queue = queue[1:]
	}
	return 0
}

func buildGraph(start string, wordList []string) map[string][]string {
	graph := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		for j := i; j < len(wordList); j++ {
			if adjacent(wordList[i], wordList[j]) {
				graph[wordList[i]] = append(graph[wordList[i]], wordList[j])
				graph[wordList[j]] = append(graph[wordList[j]], wordList[i])
			}
		}
	}

	for i := 0; i < len(wordList); i++ {
		if adjacent(start, wordList[i]) {
			graph[wordList[i]] = append(graph[wordList[i]], start)
			graph[start] = append(graph[start], wordList[i])
		}
	}
	return graph
}

func adjacent(source, target string) bool {
	var diff int
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			diff += 1
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

func main() {
	beginWord, endWord := "hit", "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	//beginWord, endWord := "hit", "cog"
	//wordList := []string{"hot","dot","dog","lot","log"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
