package main

import "fmt"
type LRUCache struct {

}


func Constructor(capacity int) LRUCache {

}


func (this *LRUCache) Get(key int) int {

}


func (this *LRUCache) Put(key int, value int)  {

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
