package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subsets(nums []int) [][]int {
	subsetsMap := recursiveSubsets(nums)
	subsetsArr := make([][]int, 0)
	for key, _ := range subsetsMap {
		arr := keyToArr(key)
		subsetsArr = append(subsetsArr, arr)
	}
	// fmt.Println(subsetsArr)
	return subsetsArr
}

func recursiveSubsets(nums []int) map[string]bool {
	subsetMap := make(map[string]bool)

	if len(nums) == 0 {
		subsetMap[""] = true
		return subsetMap
	}

	for key, _ := range recursiveSubsets(nums[1:]) {
		newKey := key + fmt.Sprintf("%d,", nums[0])
		subsetMap[newKey] = true
		subsetMap[key] = true
	}

	// fmt.Println("old map:", oldMap)
	// fmt.Println("new map:", subsetMap)

	return subsetMap
}

func keyToArr(key string) []int {
	newArr := strings.Split(key, ",")
	intArr := make([]int, 0)
	for _, item := range newArr {
		if item == "" {
			continue
		}
		intVal, _ := strconv.Atoi(item)
		intArr = append(intArr, intVal)
	}
	// fmt.Println(key, newArr, intArr)
	return intArr
}

func main() {

}
