package main

import (
	"fmt"
	"maps"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	remaining, seen := make(map[string]int), make(map[string]int)
	for _, c := range s1 {
		remaining[string(c)] += 1
	}
	maps.Copy(seen, remaining)

	//fmt.Println(remaining)
	left := 0
	for right := 0; right < len(s2); right++ {
		if _, ok := seen[string(s2[right])]; ok {
			seen[string(s2[right])] -= 1
		}
		if _, ok := remaining[string(s2[right])]; ok {
			remaining[string(s2[right])] -= 1
			if remaining[string(s2[right])] == 0 {
				delete(remaining, string(s2[right]))
			}
		}
		if len(remaining) == 0 {
			return true
		}

		if left <= right-len(s1)+1 {
			if _, ok := seen[string(s2[left])]; ok {
				seen[string(s2[left])] += 1
				if seen[string(s2[left])] > 0 {
					remaining[string(s2[left])] = seen[string(s2[left])]
					//fmt.Println("updated!")
				}
			}
			//fmt.Println("went in: leftI:", left, "remaining", string(s2[left]), " :", remaining[string(s2[left])])
			left += 1
		}
		//fmt.Println(s2[left:right+1], remaining)

	}

	return false
}

func main() {
	//s1, s2 := "ab", "eidboaoo"
	//s1, s2 := "ab", "eidbaoo"
	//s1, s2 := "adc", "dcda"
	s1, s2 := "trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine"
	fmt.Println(checkInclusion(s1, s2))
}
