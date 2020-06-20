package heapTag

import (
	"container/heap"
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
347. 前 K 个高频元素 TODO
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

 

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
 

提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	list := pairs(make([]pair, 0, k))
	for key, val := range m {
		heap.Push(&list, pair{first: val, second: key})
		if list.Len() > k {
			heap.Pop(&list)
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = (list)[i].second
	}
	return result
}

type pair struct {
	first  int
	second int
}

type pairs []pair

func (p *pairs) Len() int {
	return len(*p)
}

func (p *pairs) Less(i, j int) bool {
	return (*p)[i].first < (*p)[j].first
}

func (p *pairs) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *pairs) Push(x interface{}) {
	*p = append(*p, x.(pair))
}

func (p *pairs) Pop() interface{} {
	*p = (*p)[:len(*p)-1]
	return nil
}

func TesttopKFrequent() {
	tests := []struct {
		nums []int
		k    int
		res  []int
	}{
		{
			[]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3},
			3,
			[]int{3, 7, 10},
		},
		{
			[]int{5, 2, 5, 3, 5, 3, 1, 1, 3},
			2,
			[]int{3, 5},
		},
		{
			[]int{1, 1, 1, 1, 1, 2, 2, 3, 3, 3},
			2,
			[]int{1, 3},
		},
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
		},
	}
	for index, test := range tests {
		res := topKFrequent(test.nums, test.k)
		if !common.IntEqualSort(res, test.res) {
			fmt.Println(index, res, test.res)
		}
	}
}
