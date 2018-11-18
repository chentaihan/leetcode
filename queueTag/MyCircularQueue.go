package queueTag

import "fmt"

/**
622. 设计循环队列
 */

type MyCircularQueue struct {
	head   int
	tail   int
	size   int
	buffer []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func NewMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		head:   -1,
		tail:   -1,
		size:   k - 1,
		buffer: make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.tail == -1 {
		this.head, this.tail = 0, 0
	} else {
		if this.IsFull() {
			return false
		}
		if this.tail == this.size {
			this.tail = 0
		} else {
			this.tail++
		}
	}
	this.buffer[this.tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
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
	if this.head == this.size {
		this.head = 0
	} else {
		this.head++
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.tail+1 == this.head || (this.tail == this.size && this.head == 0)
}

func TestMyCircularQueue() {
	k := 10
	q := NewMyCircularQueue(k)
	for i := 0; i < k; i++ {
		q.EnQueue(i)
	}
	fmt.Println(q.IsFull())
	for i := 0; i < k; i++ {
		q.DeQueue()
	}
	fmt.Println(q.IsEmpty())
	for i := 0; i < k; i++ {
		q.EnQueue(i)
	}
	fmt.Println(q.IsFull())
	for i := 0; i < k; i++ {
		q.DeQueue()
	}
	fmt.Println(q.IsEmpty())
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	fmt.Println(!q.IsEmpty())
	fmt.Println(!q.IsFull())
	for !q.IsEmpty() {
		q.DeQueue()
	}
	q.EnQueue(1)
	fmt.Println(q.Front() == 1)
	fmt.Println(q.Rear() == 1)
	q.EnQueue(2)
	fmt.Println(q.Front() == 1)
	fmt.Println(q.Rear() == 2)
}
