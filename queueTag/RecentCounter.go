package queueTag

import "fmt"

/**
933. 最近的请求次数
写一个 RecentCounter 类来计算最近的请求。

它只有一个方法：ping(int t)，其中 t 代表以毫秒为单位的某个时间。

返回从 3000 毫秒前到现在的 ping 数。

任何处于 [t - 3000, t] 时间范围之内的 ping 都将会被计算在内，包括当前（指 t 时刻）的 ping。

保证每次对 ping 的调用都使用比之前更大的 t 值。

示例：

输入：inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]
输出：[null,1,2,3,3]

提示：

每个测试用例最多调用 10000 次 ping。
每个测试用例会使用严格递增的 t 值来调用 ping。
每次调用 ping 都有 1 <= t <= 10^9。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-recent-calls
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type QueueItem struct {
	Val  int
	Next *QueueItem
}

type RecentCounter struct {
	Head  *QueueItem
	Tail  *QueueItem
	queue queue
	size  int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (q *RecentCounter) Ping(t int) int {
	if q.Head == nil {
		q.Head = q.queue.Pop(t)
		q.Tail = q.Head
	} else {
		q.Tail.Next = q.queue.Pop(t)
		q.Tail = q.Tail.Next
	}
	q.size++
	count := 0
	cur := q.Head
	t -= 3000
	for cur != nil {
		if cur.Val < t {
			q.queue.Push(cur)
			cur = cur.Next
			count++
		} else {
			break
		}
	}
	q.Head = cur
	q.size -= count
	return q.size
}

type queue struct {
	Head *QueueItem
	Tail *QueueItem
}

func (q *queue) Push(item *QueueItem) {
	if q.Head == nil {
		q.Head, q.Tail = item, item
	} else {
		q.Tail.Next = item
		q.Tail = q.Tail.Next
	}
}

func (q *queue) Pop(val int) *QueueItem {
	if q.Head == nil {
		return &QueueItem{
			Val: val,
		}
	} else {
		item := q.Head
		item.Val = val
		if q.Head == q.Tail {
			q.Head, q.Tail = nil, nil
		} else {
			q.Head = q.Head.Next
		}
		return item
	}
}

func TestRecentCounter() {
	tests := []struct {
		nums  []int
		count []int
	}{
		{
			[]int{1, 2, 100, 3001, 3002, 3003, 4000, 6006, 9000, 20000},
			[]int{1, 2, 3, 4, 4, 4, 4, 2, 2, 1},
		},
	}
	for _, test := range tests {
		rc := Constructor()
		for index, t := range test.nums {
			count := rc.Ping(t)
			fmt.Println(count == test.count[index])
		}
	}
}
