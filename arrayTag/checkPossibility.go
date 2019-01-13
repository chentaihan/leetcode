package arrayTag

import "fmt"

/**
665. 非递减数列
给定一个长度为 n 的整数数组，你的任务是判断在最多改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (1 <= i < n)，满足 array[i] <= array[i + 1]。

示例 1:

输入: [4,2,3]
输出: True
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。

 */

func checkPossibility(nums []int) bool {
	count := 0
	l := len(nums) - 1
	for i := 0; i < l; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i >= 1 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return count <= 1
}

func TestcheckPossibility() {
	tests := []struct {
		nums []int
		ret  bool
	}{
		{
			[]int{-1, 4, 2, 3},
			true,
		},
		{
			[]int{2, 3, 3, 3, 4},
			true,
		},
		{
			[]int{3, 4, 2, 3},
			false,
		},
		{
			[]int{2, 3, 3, 5, 4},
			true,
		},

		{
			[]int{2, 3, 3, 2, 4},
			true,
		},
		{
			[]int{2, 3, 3, 4, 4, 1, 2},
			false,
		},
		{
			[]int{2, 3, 3, 4, 4, 1, 5, 4},
			false,
		},
		{
			[]int{2, 3, 4, 4, 4, 4, 5, 4},
			true,
		},
		{
			[]int{2, 3, 4, 5, 4, 4, 5, 4},
			false,
		},
	}

	for _, test := range tests {
		fmt.Println(checkPossibility(test.nums) == test.ret)
	}
}
