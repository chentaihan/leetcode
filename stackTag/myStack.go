package stackTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
225. 用队列实现栈
 */

type MyStack struct {
	q1 common.Queue
	q2 common.Queue
}

/** Initialize your data structure here. */
func NewMyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for !this.q1.Empty() {
		val := this.q1.Pop()
		if this.q1.Empty() {
			return val
		}
		this.q2.Push(val)
	}

	for !this.q2.Empty() {
		val := this.q2.Pop()
		if this.q2.Empty() {
			return val
		}
		this.q1.Push(val)
	}
	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for !this.q1.Empty() {
		val := this.q1.Pop()
		this.q2.Push(val)
		if this.q1.Empty() {
			return val
		}
	}

	for !this.q2.Empty() {
		val := this.q2.Pop()
		this.q1.Push(val)
		if this.q2.Empty() {
			return val
		}
	}
	return -1
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.Empty() && this.q2.Empty()
}

///**
// * Your MyStack object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * param_2 := obj.Pop();
// * param_3 := obj.Top();
// * param_4 := obj.Empty();
// */

func TestStack() {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0,},
			[]int{0},
		},
		{
			[]int{0, 1},
			[]int{1, 0},
		},
		{
			[]int{0, 1, 2},
			[]int{2, 1, 0},
		},
	}

	for _, test := range tests {
		s := NewMyStack()
		for _, item := range test.input {
			s.Push(item)
		}
		output := []int{}
		for !s.Empty() {
			output = append(output, s.Pop())
		}
		fmt.Println(common.IntEqual(test.output, output))
	}
}
