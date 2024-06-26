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
	cache := LRUCache{
		cache: make(map[int]*cacheNode),
		cap:   capacity,
	}
	first, last := &cacheNode{}, &cacheNode{}
	first.next = last
	last.prev = first
	cache.first, cache.last = first, last

	return cache
}

func (cache LRUCache) String() string {
	str := ""
	for cache.first != nil {
		str += fmt.Sprintf("%d:%d -> ", cache.first.key, cache.first.value)
		cache.first = cache.first.next
	}
	return str
}

func (cache *LRUCache) remove(node *cacheNode) {
	//fmt.Println("removing...")
	//fmt.Println(cache)

	node.prev.next = node.next
	node.next.prev = node.prev

}

func (cache *LRUCache) append(node *cacheNode) {
	node.prev = cache.last.prev
	node.next = cache.last
	cache.last.prev.next = node
	cache.last.prev = node
	//fmt.Println(cache)

}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.cache[key]; ok {
		cache.remove(node)
		cache.append(node)
		return node.value
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.cache[key]; ok {
		node.value = value
		cache.remove(node)
		cache.append(node)
	} else {
		if len(cache.cache) >= cache.cap {
			//fmt.Println("overflow")
			delete(cache.cache, cache.first.next.key)
			cache.remove(cache.first.next)

		}
		newNode := &cacheNode{
			key:   key,
			value: value,
			next:  nil,
			prev:  nil,
		}

		cache.append(newNode)
		cache.cache[key] = newNode
	}
}

func main() {
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
