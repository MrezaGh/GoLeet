package main

import (
	"fmt"
	"slices"
)

func findItinerary(tickets [][]string) []string {
	graph := make(map[string]map[string]int)
	visited := make(map[string]map[string]int)

	for _, ticket := range tickets {
		if _, ok := graph[ticket[0]]; ok {
			graph[ticket[0]][ticket[1]] += 1
			visited[ticket[0]][ticket[1]] += 1
		} else {
			graph[ticket[0]] = map[string]int{ticket[1]: 1}
			visited[ticket[0]] = map[string]int{ticket[1]: 1}
		}
	}
	//fmt.Println(graph)
	for _, neighbor := range sortedMap(graph["JFK"]) {
		if count, ok := visited["JFK"][neighbor]; ok && count >= 1 {
			visited["JFK"][neighbor] -= 1
			if path, found := dfs(neighbor, visited, len(tickets)-1); found {
				path = append(path, "JFK")
				slices.Reverse(path)
				return path
			}
			visited["JFK"][neighbor] += 1
		}
	}
	//sorted

	return nil
}

func sortedMap(neighbors map[string]int) []string {
	keys := make([]string, 0)
	for k, v := range neighbors {
		if v >= 1 {
			keys = append(keys, k)
		}
	}

	slices.Sort(keys)
	return keys
}

func dfs(node string, visited map[string]map[string]int, remainingLength int) (path []string, found bool) {
	if remainingLength == 0 {
		return []string{node}, true
	}
	for _, neighbor := range sortedMap(visited[node]) {
		visited[node][neighbor] -= 1
		if path, found := dfs(neighbor, visited, remainingLength-1); found {
			path = append(path, node)
			return path, true
		}
		visited[node][neighbor] += 1
	}

	return nil, false
}

func main() {
	tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	//tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(findItinerary(tickets))
}
