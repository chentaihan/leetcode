package common

type CircleQueue struct {
	head   int
	tail   int
	size   int
	buffer []int
}

func NewCircleQueue(cap int) *CircleQueue {
	if cap < 4 {
		cap = 4
	}
	return &CircleQueue{
		head:   -1,
		tail:   -1,
		size:   int(cap),
		buffer: make([]int, cap),
	}
}

func (queue *CircleQueue) Push(val int) bool {
	if queue.head == -1 {
		queue.head++
		queue.tail++
		queue.buffer[queue.head%queue.size] = val
	}
	return true
}
