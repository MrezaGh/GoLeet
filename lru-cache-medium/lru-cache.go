package main

import "fmt"

type cacheNode struct {
	key, value int
	next       *cacheNode
	prev       *cacheNode
}

type LRUCache struct {
	cache       map[int]*cacheNode
	first, last *cacheNode
	cap         int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*cacheNode), cap: capacity}
}

func (cache LRUCache) String() string {
	str := ""
	for cache.first != nil {
		str += fmt.Sprintf("%d:%d -> ", cache.first.key, cache.first.value)
		cache.first = cache.first.next
	}
	//str += "\n"
	return str
}

func (cache *LRUCache) refreshNode(node *cacheNode) {
	//if node.value == 2 {
	//	fmt.Println("====")
	//}
	//fmt.Println("refreshing...")
	//fmt.Println(cache)
	if node.prev != nil && node.next != nil {
		node.prev.next = node.next
		node.next.prev = node.prev

	} else if node.prev != nil && node.next == nil {
		node.prev.next = node.next
		cache.last = node.prev // will be updated later
	} else if node.prev == nil && node.next != nil {
		cache.first = node.next
		node.next.prev = nil
	} else if node.prev == nil && node.next == nil {
		cache.last = nil
		cache.first = nil
	}
	cache.appendNode(node)
	//fmt.Println(cache)
}

func (cache *LRUCache) appendNode(node *cacheNode) {
	node.next = nil
	if cache.last == nil {
		cache.last = node
	} else {
		cache.last.next = node
		node.prev = cache.last
		cache.last = cache.last.next
		//FIXME: one line or not
	}
	if cache.first == nil {
		cache.first = node
		node.prev = nil

	}

}

func (cache *LRUCache) removeOldest() {
	fmt.Println("deleted key from map:", cache.first.key, "map_value:", cache.cache[cache.first.key])
	delete(cache.cache, cache.first.key)
	if cache.last == cache.first {
		cache.last = nil
	}

	cache.first = cache.first.next
	if cache.first != nil {
		cache.first.prev = nil
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.cache[key]; ok {
		cache.refreshNode(node)
		return node.value
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.cache[key]; ok {
		node.value = value
		cache.refreshNode(node)
	} else {
		if len(cache.cache) >= cache.cap {
			cache.removeOldest()
		}
		newNode := &cacheNode{
			key:   key,
			value: value,
			next:  nil,
			prev:  nil,
		}

		cache.appendNode(newNode)
		cache.cache[key] = newNode
	}
}

func main() {
	//capacity := 2
	//cache := Constructor(capacity)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(1))
	//cache.Put(3, 3)
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 4)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(4))

	//capacity := 2
	//cache := Constructor(capacity)
	//cache.Put(2, 1)
	//cache.Put(3, 2)
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 3)
	//fmt.Println(cache.Get(2))
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(4))

	capacity := 3
	cache := Constructor(capacity)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))
	cache.Put(5, 5)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(5))

}
