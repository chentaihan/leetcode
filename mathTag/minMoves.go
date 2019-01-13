package mathTag

import "fmt"

/**
453. 最小移动次数使数组元素相等
给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。

示例:

输入:
[1,2,3]

输出:
3

解释:
只需要3次移动（注意每次移动会增加两个元素的值）：
[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
分析：每次移动使n-1个元素加一，直到最大值，倒过来就是每次可使一个元素减一，直到最小值
 */

func minMoves(nums []int) int {
	min := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}

func TestminMoves() {
	tests := []struct {
		nums  []int
		count int
	}{
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{2, 2},
			0,
		},
		{
			[]int{1, 3},
			2,
		},
		{
			[]int{1, 7},
			6,
		},
		{
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{2, 2, 3},
			1,
		},
		{
			[]int{2, 2, 6},
			4,
		},
		{
			[]int{1, 2, 3, 4},
			6,
		},
	}

	for _, test := range tests {
		count := minMoves(test.nums)
		fmt.Println(count == test.count)
	}
}
