package common

import "fmt"

type Stack struct {
	buffer []int
	size   int
}

/** Initialize your data structure here. */
func NewStack() Stack {
	return Stack{}
}

/** Push element x onto stack. */
func (this *Stack) Push(x int) {
	if len(this.buffer) == this.size {
		this.buffer = append(this.buffer, x)
	} else {
		this.buffer[this.size] = x
	}
	this.size++
}

/** Removes the element on top of the stack and returns that element. */
func (this *Stack) Pop() int {
	if this.size == 0 {
		return -1
	}
	val := this.buffer[this.size-1]
	this.size--
	return val
}

/** Get the top element. */
func (this *Stack) Top() int {
	if this.size == 0 {
		return -1
	}
	return this.buffer[this.size-1]
}

/** Returns whether the stack is empty. */
func (this *Stack) Empty() bool {
	return this.size == 0
}

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
		s := NewStack()
		for _, item := range test.input {
			s.Push(item)
		}
		output := []int{}
		for !s.Empty() {
			output = append(output, s.Pop())
		}
		fmt.Println(IntEqual(test.output, output))
	}
}
