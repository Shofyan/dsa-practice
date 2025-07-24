package main

import "fmt"

// LRUCache is a Least Recently Used cache.
type LRUCache struct {
	capacity int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

// DLinkedNode is a node in a doubly linked list.
type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

// Constructor initializes an LRUCache with a given capacity.
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get returns the value of the key if the key exists, otherwise returns -1.
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

// Put sets or inserts the value if the key is not already present.
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.moveToHead(node)
	} else {
		newNode := &DLinkedNode{key: key, value: value}
		lru.cache[key] = newNode
		lru.addNode(newNode)

		if len(lru.cache) > lru.capacity {
			tail := lru.popTail()
			delete(lru.cache, tail.key)
		}
	}
}

// addNode adds a new node right after head.
func (lru *LRUCache) addNode(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode removes an existing node from the linked list.
func (lru *LRUCache) removeNode(node *DLinkedNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// moveToHead moves a certain node in the linked list to the head.
func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addNode(node)
}

// popTail removes the last node from the tail of the linked list.
func (lru *LRUCache) popTail() *DLinkedNode {
	res := lru.tail.prev
	lru.removeNode(res)
	return res
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Printf("Get(1): %d
", lru.Get(1)) // returns 1
	lru.Put(3, 3)                            // evicts key 2
	fmt.Printf("Get(2): %d
", lru.Get(2)) // returns -1 (not found)
	lru.Put(4, 4)                            // evicts key 1
	fmt.Printf("Get(1): %d
", lru.Get(1)) // returns -1 (not found)
	fmt.Printf("Get(3): %d
", lru.Get(3)) // returns 3
	fmt.Printf("Get(4): %d
", lru.Get(4)) // returns 4
}
