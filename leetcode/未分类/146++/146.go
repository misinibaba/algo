package main

import "fmt"
type LRUCache struct {
	cache map[int]*linkNode
	capacity int
	size int
	head, tail *linkNode
}

type linkNode struct {
	key, val int
	prev, next *linkNode
}

func newLinkNode(key, val int) *linkNode {
	return &linkNode{
		key:  key,
		val:  val,
		prev: nil,
		next: nil,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*linkNode),
		capacity: capacity,
		size:     0,
		head:     newLinkNode(0, 0),
		tail:     newLinkNode(0, 0),
	}
	l.tail.prev = l.head
	l.head.next = l.tail
	return l
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		newNode := newLinkNode(key, value)
		this.addToHead(newNode)
		this.cache[key] = newNode
		this.size++
		if this.size > this.capacity {
			rmNode := this.removeTail()
			delete(this.cache, rmNode.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.val = value
		this.moveToHead(node)
	}
	fmt.Println(this.cache)
}

func (this *LRUCache) addToHead(node *linkNode)  {
	node.next = this.head.next
	node.prev = this.head
	node.next.prev = node
	node.prev.next = node
}

func (this *LRUCache) moveToHead(node *linkNode)  {
	node.prev.next = node.next
	node.next.prev = node.prev
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *linkNode)  {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *linkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func main() {
	res := 0
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	res = lRUCache.Get(1)    // 返回 1
	fmt.Println(res)
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	res = lRUCache.Get(2)    // 返回 -1 (未找到)
	fmt.Println(res)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	res = lRUCache.Get(1)    // 返回 -1 (未找到)
	fmt.Println(res)
	res = lRUCache.Get(3)    // 返回 3
	fmt.Println(res)
	res = lRUCache.Get(4)    // 返回 4
	fmt.Println(res)
}
