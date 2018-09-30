package treeTag

type queueItem struct {
	next *queueItem
	val  *TreeNode
}

type Queue struct {
	head *queueItem
	tail *queueItem
}

func (s *Queue) Push(val *TreeNode) {
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

func (s *Queue) Pop() *TreeNode {
	if s.head != nil {
		val := s.head.val
		s.head = s.head.next
		if s.head == nil {
			s.tail = nil
		}
		return val
	}
	return nil
}

func (s *Queue) Empty() bool {
	return  s.head == nil
}

func (s *Queue) Peek() *TreeNode {
	if s.head != nil {
		return s.head.val
	}
	return nil
}


type queueItemEx struct {
	next *queueItemEx
	val  *TreeNode
	level int
}

type QueueEx struct {
	head *queueItemEx
	tail *queueItemEx
}

func (s *QueueEx) Push(val *TreeNode, level int) {
	item := &queueItemEx{
		val: val,
		level:level,
	}
	if s.tail == nil {
		s.head = item
	} else {
		s.tail.next = item
	}
	s.tail = item
}

func (s *QueueEx) Pop() (*TreeNode, int) {
	if s.head != nil {
		val := s.head.val
		level := s.head.level
		s.head = s.head.next
		if s.head == nil {
			s.tail = nil
		}
		return val, level
	}
	return nil, -1
}

func (s *QueueEx) Empty() bool {
	return  s.head == nil
}

func (s *QueueEx) Peek() *TreeNode {
	if s.head != nil {
		return s.head.val
	}
	return nil
}
