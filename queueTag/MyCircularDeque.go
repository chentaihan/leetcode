package queueTag

import "fmt"

/**
641. 设计循环双端队列
 */

type MyCircularDeque struct {
	head   int
	tail   int
	size   int
	buffer []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func NewMyCircularDeque(k int) MyCircularDeque {
	return MyCircularDeque{
		head:   -1,
		tail:   -1,
		size:   k,
		buffer: make([]int, k),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	head := this.head
	if head == -1 {
		this.head, this.tail = 0, 0
	} else {
		if this.IsFull() {
			return false
		}
		if head == 0 {
			head = this.size - 1
		} else {
			head--
		}
		this.head = head
	}
	this.buffer[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	tail := this.tail
	if tail == -1 {
		this.head, this.tail = 0, 0
	} else {
		if this.IsFull() {
			return false
		}
		if tail == this.size-1 {
			tail = 0
		} else {
			tail++
		}
		this.tail = tail
	}
	this.buffer[this.tail] = value
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	//0个元素
	if this.IsEmpty() {
		return false
	}
	//1个元素
	if this.head == this.tail {
		this.head, this.tail = -1, -1
		return true
	}
	//多个元素
	if this.head == this.size-1 {
		this.head = 0
	} else {
		this.head++
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	//0个元素
	if this.IsEmpty() {
		return false
	}
	//1个元素
	if this.head == this.tail {
		this.head, this.tail = -1, -1
		return true
	}
	//多个元素
	if this.tail == 0 {
		this.tail = this.size - 1
	} else {
		this.tail--
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.tail]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == -1
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.tail + 1 == this.head || (this.tail  == this.size -1 && this.head == 0)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := NewMyCircularQueue(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func TestMyCircularDeque() {
	k := 10
	q := NewMyCircularDeque(k)
	for i := 0; i < k; i++ {
		q.InsertLast(i)
	}
	fmt.Println(q.IsFull())
	for i := 0; i < k; i++ {
		q.DeleteLast()
	}
	fmt.Println(q.IsEmpty())
	for i := 0; i < k; i++ {
		q.InsertFront(i)
	}
	fmt.Println(q.IsFull())
	for i := 0; i < k; i++ {
		q.DeleteFront()
	}
	fmt.Println(q.IsEmpty())
	for i := 0; i < 5; i++ {
		q.InsertFront(i)
	}
	fmt.Println(!q.IsEmpty())
	fmt.Println(!q.IsFull())
	for !q.IsEmpty() {
		q.DeleteFront()
	}
	q.InsertFront(1)
	fmt.Println(q.GetFront() == 1)
	fmt.Println(q.GetRear() == 1)
	q.InsertLast(2)
	fmt.Println(q.GetFront() == 1)
	fmt.Println(q.GetRear() == 2)
}
