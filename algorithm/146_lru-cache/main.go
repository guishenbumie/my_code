package main

import "fmt"

/* LRU 缓存 */

//请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
//示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//提示：
//
//1 <= capacity <= 3000
//0 <= key <= 10000
//0 <= value <= 105
//最多调用 2 * 105 次 get 和 put

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
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

// 头尾两个点使用虚拟节点
func constructList() *List {
	h := &Node{Key: 0, Val: 0}
	t := &Node{Key: 0, Val: 0}
	h.Next = t
	t.Prev = h
	return &List{head: h, tail: t, size: 0}
}

// 在链表尾部增加数据
func (l *List) addLast(n *Node) {
	n.Prev = l.tail.Prev
	n.Next = l.tail
	l.tail.Prev.Next = n
	l.tail.Prev = n
	l.size++
}

// 删除链表中指定的节点
func (l *List) remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	l.size--
}

// 删除链表中第一个节点，并返回该节点
func (l *List) removeFirst() *Node {
	if l.head.Next == l.tail {
		return nil
	}
	n := l.head.Next
	l.remove(n)
	return n
}

func (l *List) getSize() int {
	return l.size
}

// 使用双向链表+哈希表实现
type LRUCache struct {
	m   map[int]*Node
	l   *List
	cap int //最大容量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:   make(map[int]*Node),
		l:   constructList(),
		cap: capacity,
	}
}

// 将某个key提升为最近使用的
func (this *LRUCache) makeRecently(key int) {
	n := this.m[key]
	//先从链表删除，再添加到尾部
	this.l.remove(n)
	this.l.addLast(n)
}

// 添加最近使用的节点
func (this *LRUCache) addRecently(key, val int) {
	n := &Node{Key: key, Val: val}
	//新增元素，链表和map都要增加
	this.l.addLast(n)
	this.m[key] = n
}

// 删除某一个key
func (this *LRUCache) deleteKey(key int) {
	n := this.m[key]
	//删除元素，链表和map都要删除
	this.l.remove(n)
	delete(this.m, key)
}

// 删除最近最少使用的元素
func (this *LRUCache) removeLeastRecently() {
	//删除元素，链表和map都要删除
	n := this.l.removeFirst()
	delete(this.m, n.Key)
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}
	//提升为最近使用
	this.makeRecently(key)
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; ok { //已经存在
		this.deleteKey(key)
	} else {
		if this.cap == len(this.m) {
			this.removeLeastRecently()
		}
	}
	this.addRecently(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
