package stackTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
232. 用栈实现队列
 */

type MyQueue struct {
	s1 common.Stack
	s2 common.Stack
}

/** Initialize your data structure here. */
func NewMyQueue() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for !this.s1.Empty() {
		this.s2.Push(this.s1.Pop())
	}
	return this.s2.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for !this.s1.Empty() {
		this.s2.Push(this.s1.Pop())
	}
	return this.s2.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.Empty() && this.s2.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

 func TestMyQeue(){
	 tests := []struct {
		 input  []int
		 output []int
	 }{
		 {
			 []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			 []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		 },
		 {
			 []int{0,},
			 []int{0},
		 },
		 {
			 []int{0, 1},
			 []int{0, 1},
		 },
		 {
			 []int{0, 1,2},
			 []int{0, 1,2},
		 },
	 }

	 for _, test := range tests {
		 s := NewMyQueue()
		 for _, item := range test.input {
			 s.Push(item)
		 }
		 output := []int{}
		 for !s.Empty(){
			 output = append(output, s.Pop())
		 }
		 fmt.Println(common.IntEqual(test.output, output))
	 }
 }