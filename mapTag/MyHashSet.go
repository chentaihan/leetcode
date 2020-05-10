package mapTag

/**
705. 设计哈希集合
不使用任何内建的哈希表库设计一个哈希集合

具体地说，你的设计应该包含以下的功能

add(value)：向哈希集合中插入一个值。
contains(value) ：返回哈希集合中是否存在这个值。
remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

示例:

MyHashSet hashSet = new MyHashSet();
hashSet.add(1);        
hashSet.add(2);        
hashSet.contains(1);    // 返回 true
hashSet.contains(3);    // 返回 false (未找到)
hashSet.add(2);          
hashSet.contains(2);    // 返回 true
hashSet.remove(2);          
hashSet.contains(2);    // 返回  false (已经被删除)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-hashset
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type setItem struct {
	val  int
	next *setItem
}

type MyHashSet struct {
	buf []*setItem
}

/** Initialize your data structure here. */
func Constructor1() MyHashSet {
	return MyHashSet{
		buf: make([]*setItem, 128),
	}
}

func (this *MyHashSet) Add(key int) {
	node := this.buf[key&127]
	for item := node; item != nil; item = item.next {
		if item.val == key {
			return
		}
	}
	this.buf[key&127] = &setItem{
		val:  key,
		next: node,
	}
}

func (this *MyHashSet) Remove(key int) {
	node := this.buf[key&127]
	if node == nil {
		return
	}
	if node.val == key {
		this.buf[key&127] = node.next
		return
	}
	for item := node; item != nil; item = item.next {
		if item.next != nil && item.next.val == key {
			item.next = item.next.next
			break
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	node := this.buf[key&127]
	for item := node; item != nil; item = item.next {
		if item.val == key {
			return true
		}
	}
	return false
}
