package main

import "fmt"

type TimeMap struct {
	m map[string][]entry
}

type entry struct {
	val       string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]entry)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.m[key] = append(tm.m[key], entry{val: value, timestamp: timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	arr, ok := tm.m[key]
	if !ok {
		return ""
	}
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) / 2
		if arr[mid].timestamp < timestamp {
			left = mid + 1
		} else if arr[mid].timestamp >= timestamp {
			right = mid
		}
	}
	if arr[right].timestamp == timestamp {
		return arr[right].val
	} else if right > 0 {
		if timestamp < arr[right].timestamp {
			return arr[right-1].val
		} else {
			return arr[right].val
		}
	} else if timestamp > arr[right].timestamp {
		return arr[right].val
	} else if timestamp < arr[right].timestamp {
		return ""
	}
	return "fff"
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	tm := Constructor()
	//tm.Set("love", "high", 10)
	//tm.Set("love", "low", 20)
	//fmt.Println(tm.Get("love", 5))
	//fmt.Println(tm.Get("love", 10))
	//fmt.Println(tm.Get("love", 15))
	//fmt.Println(tm.Get("love", 20))
	//fmt.Println(tm.Get("love", 25))

	//tm.Set("foo", "bar", 1)
	//fmt.Println(tm.Get("foo", 1))
	//fmt.Println(tm.Get("foo", 3))
	//tm.Set("foo", "bar2", 4)
	//fmt.Println(tm.Get("foo", 4))
	//fmt.Println(tm.Get("foo", 5))

	tm.Set("a", "bar", 1)
	tm.Set("x", "b", 3)
	fmt.Println(tm.Get("b", 3))
	tm.Set("foo", "bar2", 4)
	fmt.Println(tm.Get("foo", 4))
	fmt.Println(tm.Get("foo", 5))

}
