package title

import "fmt"

/**
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：

输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000。
https://leetcode-cn.com/problems/max-consecutive-ones/description/
 */

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}
	if count > maxCount {
		return count
	}
	return maxCount
}

func TestFindMaxConsecutiveOnes(){
	tests := []struct{
		nums []int
		count int
	}{
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{0,0,0,0,0,0,0,0,0,0,0,0},
			0,
		},
		{
			[]int{0,0,0,1,0,0,0,0,0,0,0,0},
			1,
		},
		{
			[]int{1,1},
			2,
		},
		{
			[]int{1,1,1},
			3,
		},
		{
			[]int{1,1,1,0,1,1},
			3,
		},
		{
			[]int{1,1,0,1,1,1},
			3,
		},
		{
			[]int{1,1,1,1,1,1,1,1,1,1},
			10,
		},
	}
	for _,test := range tests{
		fmt.Println(findMaxConsecutiveOnes(test.nums) == test.count)
	}
}
