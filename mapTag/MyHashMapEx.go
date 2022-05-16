package mapTag

import "fmt"

/**
706. 设计哈希映射
不使用任何内建的哈希表库设计一个哈希映射

具体地说，你的设计应该包含以下的功能

put(key, value)：向哈希映射中插入(键,值)的数值对。如果键对应的值已经存在，更新这个值。
get(key)：返回给定的键所对应的值，如果映射中不包含这个键，返回-1。
remove(key)：如果映射中存在这个键，删除这个数值对。

示例：

MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // 返回 1
hashMap.get(3);            // 返回 -1 (未找到)
hashMap.put(2, 1);         // 更新已有的值
hashMap.get(2);            // 返回 1
hashMap.remove(2);         // 删除键为2的数据
hashMap.get(2);            // 返回 -1 (未找到)

注意：

所有的值都在 [1, 1000000]的范围内。
操作的总数目在[1, 10000]范围内。
不要使用内建的哈希库。
*/

type NodeEx struct {
	key   int
	value int
}

type MyHashMapEx struct {
	buckets    [][]NodeEx
	bucketSize int
}

/** Initialize your data structure here. */
func ConstructorEx() MyHashMapEx {
	bucketSize := 1 << 10
	return MyHashMapEx{
		buckets:    make([][]NodeEx, bucketSize),
		bucketSize: bucketSize - 1,
	}
}

/** value will always be non-negative. */
func (this *MyHashMapEx) Put(key int, value int) {
	bucketIndex := key & this.bucketSize
	nodeList := this.buckets[bucketIndex]
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].key == key {
			nodeList[i].value = value
			return
		}
	}
	this.buckets[bucketIndex] = append(this.buckets[bucketIndex], NodeEx{
		key:   key,
		value: value,
	})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMapEx) Get(key int) int {
	bucketIndex := key & this.bucketSize
	nodeList := this.buckets[bucketIndex]
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].key == key {
			return nodeList[i].value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMapEx) Remove(key int) {
	bucketIndex := key & this.bucketSize
	nodeList := this.buckets[bucketIndex]
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].key == key {
			this.buckets[bucketIndex] = append(nodeList[:i], nodeList[i+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func TestMyHashMapEx() {
	obj := ConstructorEx()
	tests := []struct {
		key   int
		value int
	}{
		{
			1, 1,
		},
		{
			2, 2,
		},
		{
			3, 3,
		},
		{
			1 << 10, 1 << 10,
		},
		{
			(1 << 10) + 1, (1 << 10) + 1,
		},
		{
			(1 << 10) + 2, (1 << 10) + 2,
		},
		{
			(1 << 10) + 3, (1 << 10) + 3,
		},
		{
			(1 << 11) + 1, (1 << 11) + 1,
		},
		{
			(1 << 11) + 2, (1 << 11) + 2,
		},
		{
			(1 << 11) + 3, (1 << 11) + 3,
		},
	}

	for _, test := range tests {
		obj.Put(test.key, test.value)
	}
	obj.Remove(1)
	obj.Remove((1 << 10) + 1)
	obj.Remove((1 << 11) + 1)
	for _, test := range tests {
		value := obj.Get(test.key)
		if value != -1 {
			fmt.Println(value == test.value)
		} else {
			fmt.Println("key= ", test.key, " not found")
		}

	}
}
