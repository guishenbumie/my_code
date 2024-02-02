package main

import (
	"container/list"
	"fmt"
)

/* LRU 缓存 */

// 和146题一样
func main() {
	lru := Constructor(2)
	fmt.Println(lru.Get(2))
	lru.Put(2, 6)
	fmt.Println(lru.Get(1))
	lru.Put(1, 5)
	lru.Put(1, 2)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
}

type Node struct {
	key int
	val int
}

type LRUCache struct {
	m   map[int]*list.Element
	l   *list.List
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:   map[int]*list.Element{},
		l:   list.New(),
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if !ok {
		return -1
	}
	this.l.MoveToFront(e)
	return e.Value.(*Node).val
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	var node *Node
	if !ok {
		node = &Node{key: key, val: value}
		if len(this.m) == this.cap {
			del := this.l.Remove(this.l.Back()).(*Node)
			delete(this.m, del.key)
		}
		e := this.l.PushFront(node)
		this.m[key] = e
	} else {
		node = e.Value.(*Node)
		node.val = value
		this.l.MoveToFront(e)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
