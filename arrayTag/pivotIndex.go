package arrayTag

import "fmt"

/**
724. 寻找数组的中心索引
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。

我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

示例 1:

输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。
示例 2:

输入:
nums = [1, 2, 3]
输出: -1
解释:
数组中不存在满足此条件的中心索引。
说明:

nums 的长度范围为 [0, 10000]。
任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-pivot-index
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

方案一：从左往右开始累积，从右往走开始累加，存在一个值，他的左右的累加和相等
方案二：1先求出总和，在从左往右求和，左边和 == 总和 - 当前值 - 左边和
*/

func pivotIndex(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	leftSum := make([]int, len(nums))
	leftSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i]
	}

	rightSum := make([]int, len(nums))
	rightSum[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i]
	}

	if rightSum[1] == 0 {
		return 0
	}
	for i := 1; i < len(nums)-1; i++ {
		if leftSum[i-1] == rightSum[i+1] {
			return i
		}
	}
	if leftSum[len(leftSum)-2] == 0 {
		return len(leftSum) - 1
	}
	return -1
}

func TestpivotIndex() {
	tests := []struct {
		nums  []int
		index int
	}{
		{
			[]int{-1, -1, -1, 0, 1, 1},
			0,
		},
		{
			[]int{-1, -1, -1, 0, 3, 1},
			5,
		},
		{
			[]int{-1, 7, -3, 7, -5, 4, -6, 6},
			2,
		},
		{
			[]int{-1, -1, -1, 0, -1, -1},
			2,
		},
		{
			[]int{-1, -1, -1, 0, -1, -1, -1},
			3,
		},
		{
			[]int{1, 7, 3, 6, 5, 6},
			3,
		},
		{
			[]int{-1, -7, -3, -6, -5, -6},
			3,
		},
		{
			[]int{0, -1, -1, -1, 0, -1, 0, -1, -1, -1},
			5,
		},
		{
			[]int{-1, -1, -1, -1, -1, 0},
			2,
		},
		{
			[]int{-1, -1, -1, -1, -1, -1},
			-1,
		},
		{
			[]int{0, -1, -1, -1, -1, -1, -1, -1},
			4,
		},
		{
			[]int{1, 2, 3},
			-1,
		},
		{
			[]int{1, 2, 1},
			1,
		},
		{
			[]int{1, 5, 2, 1, 5},
			2,
		},
		{
			[]int{10, 5, 2, 5, 1, 5, 4},
			2,
		},
		{
			[]int{},
			-1,
		},
		{
			[]int{1},
			-1,
		},
	}
	for _, test := range tests {
		index := pivotIndex(test.nums)
		fmt.Println(index == test.index)
	}
}
