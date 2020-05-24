package heapTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
面试题 17.14. 最小K个数
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

示例：

输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]
提示：

0 <= len(arr) <= 100000
0 <= k <= min(100000, len(arr))

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-k-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	_smallestK(arr, k)
	return arr[:k]
}

func _smallestK(arr []int, k int) {
	index := partion(arr)
	if index+1 == k {
		return
	} else if index < k {
		_smallestK(arr[index+1:], k-index-1)
	} else {
		_smallestK(arr[:index], k)
	}
}

func partion(arr []int) int {
	val := arr[0]
	start, end := 0, len(arr)-1
	for {
		for start < end && arr[end] > val {
			end--
		}
		for start <= end && arr[start] <= val {
			start++
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		} else {
			if start > 0 {
				arr[start-1], arr[0] = arr[0], arr[start-1]
			}
			return start - 1
		}
	}

	return start
}

func TestsmallestK() {
	tests := []struct {
		arr []int
		k   int
		res []int
	}{
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{1, 2},
			0,
			[]int{0},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
		{
			[]int{3, 1, 2},
			1,
			[]int{1},
		},
		{
			[]int{3, 1, 2, 4},
			2,
			[]int{1, 2},
		},
		{
			[]int{3, 1, 2, 4, 5},
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{3, 1, 2, 4, 5, 6},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{3, 1, 7, 2, 4, 5, 6},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{3, 8, 1, 7, 2, 4, 5, 6},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{3, 8, 1, 7, 2, 4, 5, 6},
			5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 8, 1, 7, 2, 4, 5, 6},
			6,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{3, 8, 1, 7, 2, 4, 5, 6},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{3, 8, 1, 7, 2, 4, 5, 6},
			8,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{3, 9, 8, 1, 7, 2, 4, 5, 6},
			5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 9, 8, 1, 7, 2, 4, 5, 6},
			6,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{3, 9, 8, 1, 7, 2, 4, 5, 6},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{3, 9, 8, 1, 7, 2, 4, 5, 6},
			8,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test := range tests {
		result := smallestK(test.arr, test.k)
		fmt.Println(common.IntEqualSort(result, test.res))
	}
}
