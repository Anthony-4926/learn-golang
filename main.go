package main

type node struct {
	key, val  int
	pre, next *node
}

type LRUCache struct {
	m        map[int]*node
	head     *node
	end      *node
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        map[int]*node{},
		head:     new(node),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if !ok {
		return -1
	}
	this.toFirst(v)
	return v.val
}

func (this *LRUCache) toFirst(cur *node) {
	this.deleteFromList(cur)
	this.insert(cur)
}

func (this *LRUCache) deleteFromList(cur *node) {
	// 如果删除的是一个不存在的节点，直接返回就好
	if cur.pre == nil {
		return
	}

	pre := cur.pre
	next := cur.next

	pre.next = next
	if next != nil {
		next.pre = pre
	} else {
		this.end = pre
	}

}

func (this *LRUCache) insert(cur *node) {
	cur.next = this.head.next
	if this.head.next != nil {
		this.head.next.pre = cur
	}
	this.head.next = cur
	cur.pre = this.head
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.m) == this.capacity {
		this.deleteFromList(this.end)
		delete(this.m, key)
	}
	v, ok := this.m[key]
	if ok {
		v.val = value
	}
	cur := node{
		key: key,
		val: value,
	}
	this.toFirst(&cur)
	this.m[key] = &cur
}
