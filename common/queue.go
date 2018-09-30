package common

import "fmt"

type queueItem struct {
	next *queueItem
	val  int
}

type Queue struct {
	head *queueItem
	tail *queueItem
}

func NewQueue() Queue {
	return Queue{}
}

func (s *Queue) Push(val int) {
	item := &queueItem{
		val: val,
	}
	if s.tail == nil {
		s.head = item
	} else {
		s.tail.next = item
	}
	s.tail = item
}

func (s *Queue) Pop() int {
	if s.head != nil {
		val := s.head.val
		s.head = s.head.next
		if s.head == nil {
			s.tail = nil
		}
		return val
	}
	return -1
}

func (s *Queue) Empty() bool {
	return  s.head == nil
}

func (s *Queue) Peek() int {
	if s.head != nil {
		return s.head.val
	}
	return -1
}

func TestQueue(){
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
		s := NewQueue()
		for _, item := range test.input {
			s.Push(item)
		}
		output := []int{}
		for !s.Empty(){
			output = append(output, s.Pop())
		}
		fmt.Println(IntEqual(test.output, output))
	}
}
