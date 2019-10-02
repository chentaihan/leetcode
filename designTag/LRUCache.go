package designTag

import "fmt"

//146 LRU缓存机制
//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
// 进阶:
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

type LRUCache struct {
	cap     int
	mBuffer map[int]*node
	head    *node
	tail    *node
}

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}

func Constructor(capacity int) LRUCache {
	if capacity < 0 {
		capacity = 0
	}
	mBuffer := make(map[int]*node, capacity)
	return LRUCache{
		cap:     capacity,
		mBuffer: mBuffer,
	}
}

func (this *LRUCache) Get(key int) int {
	cur, exist := this.mBuffer[key]
	if !exist {
		return -1
	}
	if this.head == cur {
		return cur.value
	}
	if this.tail == cur {
		this.tail = this.tail.prev
		if this.tail == nil {
			this.tail = this.head
		}
	}
	prev := cur.prev
	if prev != nil {
		prev.next = cur.next
	}
	if cur.next != nil {
		cur.next.prev = prev
	}
	cur.prev = nil
	cur.next = this.head
	this.head.prev = cur
	this.head = cur
	return cur.value
}

func (this *LRUCache) Put(key int, value int) {
	cur, exist := this.mBuffer[key]
	if exist {
		cur.value = value
		if this.head == cur {
			return
		}
		if this.tail == cur {
			this.tail = this.tail.prev
			if this.tail == nil {
				this.tail = this.head
			}
		}
		prev := cur.prev
		if prev != nil {
			prev.next = cur.next
		}
		if cur.next != nil {
			cur.next.prev = prev
		}
		cur.prev = nil
		cur.next = this.head
		this.head.prev = cur
	} else {
		cur = &node{
			key:   key,
			value: value,
		}
		this.mBuffer[key] = cur
		cur.next = this.head
		if this.head != nil {
			this.head.prev = cur
		} else {
			this.tail = cur
		}
		if len(this.mBuffer) > this.cap {
			if this.tail != nil {
				delete(this.mBuffer, this.tail.key)
				this.tail = this.tail.prev
				if this.tail != nil {
					this.tail.next = nil
				}
			}
		}
	}
	this.head = cur
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRUCache() {
	tests := [][]struct {
		key   int
		val   int
		isGet bool
	}{
		{
			{
				2, 1, false,
			},
			{
				2, 0, true,
			},
			{
				3, 2, false,
			},
			{
				2, 0, true,
			},
			{
				3, 0, true,
			},
		},
	}

	for _, item := range tests {
		lru := Constructor(1)
		for _, test := range item {
			if test.isGet {
				fmt.Println(lru.Get(test.key))
			} else {
				lru.Put(test.key, test.val)
				fmt.Println("null")
			}
		}
	}
}
