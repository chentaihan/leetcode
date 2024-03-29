package dynamicTag

import "fmt"

/**
300. 最长上升子序列
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	buf := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		value := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if buf[j] >= value {
					value = buf[j] + 1
				}
			}
		}
		buf[i] = value
		if value > max {
			max = value
		}
	}
	return max + 1
}

func lengthOfLISEx(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	buf := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		count := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && buf[j] >= count {
				count = buf[j] + 1
			}
		}
		buf[i] = count
		if count > max {
			max = count
		}
	}
	return max + 1
}

func TestlengthOfLIS() {
	tests := []struct {
		nums []int
		max  int
	}{
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{9, 2, 5, 3, 7, 8, 101, 18},
			5,
		},
		{
			[]int{1, 9, 2, 5, 3, 7, 8, 101, 18},
			6,
		},
		{
			[]int{1, 9, 2, 5, 3, 9, 7, 8, 101, 18},
			6,
		},
	}
	for _, test := range tests {
		max := lengthOfLISEx(test.nums)
		fmt.Println(max, test.max)
	}
}
