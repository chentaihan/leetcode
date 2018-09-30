package title

import "fmt"

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
https://leetcode-cn.com/problems/maximum-subarray/description/
 */

func maxSubArray(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	max, sum := nums[0], nums[0]
	for i := 1; i < count; i++ {
		if sum < 0 {
			sum = nums[i]
		}else{
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func TestMaxSubArray() {
	tests := []struct {
		nums []int
		sum  int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{-2, 1, -3, -4, -1, -2, -1, -5, -4},
			1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			21,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			21,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 0},
			21,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 0},
			21,
		},
		{
			[]int{0, 0, 1, 2, 3, 4, 5, 6, 0, 0},
			21,
		},
		{
			[]int{0, -1, -2, -3, -4, -5, -6},
			0,
		},
		{
			[]int{-1, -2, -3, -4, -5, -6},
			-1,
		},
		{
			[]int{-7, -1, -2, -3, -4, -5, -6},
			-1,
		},
		{
			[]int{-7, -2, -3, -4, -5, -6, -1},
			-1,
		},
		{
			[]int{-7, -2, -3, -1, -4, -5, -6},
			-1,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{-1},
			-1,
		},
		{
			[]int{0, 0, 0, 0, 0, 0},
			0,
		},
		{
			[]int{1, 0, -1},
			1,
		},
		{
			[]int{-1, 0, 1},
			1,
		},
		{
			[]int{2, -1, 0, 1},
			2,
		},
		{
			[]int{2, -1, 0, 2},
			3,
		},
		{
			[]int{-2, 1, 0, 2},
			3,
		},
		{
			[]int{-2, 1, 0, 3},
			4,
		},
		{
			[]int{3, -2, 3, -2, 0, 2},
			4,
		},
	}
	for _, test := range tests {
		fmt.Println(maxSubArray(test.nums) == test.sum)
	}
}
