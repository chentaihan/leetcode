package designTag

import "fmt"

//155最小栈
//实现：栈 + 排序双向链表
//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//
// push(x) -- 将元素 x 推入栈中。
// pop() -- 删除栈顶的元素。
// top() -- 获取栈顶元素。
// getMin() -- 检索栈中的最小元素。
//
//
// 示例:
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//

type MinStack struct {
	list []*stackNode
	head *stackNode
	tail *stackNode
}

type stackNode struct {
	value int
	prev  *stackNode
	next  *stackNode
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	node := &stackNode{
		value: x,
	}
	this.list = append(this.list, node)
	tmpNode := this.head
	if this.head == nil {
		this.head, this.tail = node, node
		return
	}
	for tmpNode != nil && tmpNode.value < x {
		tmpNode = tmpNode.next
	}
	if tmpNode == nil {
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	} else {
		node.next = tmpNode
		node.prev = tmpNode.prev
		if tmpNode.prev != nil {
			tmpNode.prev.next = node
		} else {
			this.head = node
		}
		tmpNode.prev = node
	}

}

func (this *MinStack) Pop() {
	if len(this.list) == 0 {
		return
	}
	node := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	if this.head == node {
		if this.tail == node {
			this.head, this.tail = nil, nil
			return
		}
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
	} else if this.tail == node {
		prev := this.tail.prev
		prev.next = nil
		this.tail = prev
	} else {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev
	}
}

func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return -1
	}
	return this.list[len(this.list)-1].value
}

func (this *MinStack) GetMin() int {
	if len(this.list) == 0 {
		return -1
	}
	return this.head.value
}

func TestMinStackConstructor() {
	tests := []struct {
		list []int
		min  []int
	}{
		{
			[]int{7},
			[]int{7},
		},
		{
			[]int{7, 3, 5, 7, 9, 2, 0, 4, 6, 8, 10},
			[]int{7, 3, 3, 3, 3, 2, 0, 0, 0, 0, 0},
		},
		{
			[]int{10, 8, 6, 9, 7, 5, 4, 2, 3, 1, 0},
			[]int{10, 8, 6, 6, 6, 5, 4, 2, 2, 1, 0},
		},
	}
	for _, test := range tests {
		stack := MinStackConstructor()
		for _, item := range test.list {
			stack.Push(item)
		}
		for {
			minValue := stack.GetMin()
			if minValue == -1 {
				break
			}
			stack.Pop()
			result := test.min[len(test.min)-1]
			test.min = test.min[:len(test.min)-1]
			if result != minValue {
				fmt.Println("minValue=", minValue, ", result=", result)
			} else {
				fmt.Println("ok")
			}
		}

	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
