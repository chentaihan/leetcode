package title

import (
	"sort"
	"fmt"
)

/**
给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

示例 1:

输入: [1,4,3,2]

输出: 4
解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
提示:

n 是正整数,范围在 [1, 10000].
数组中的元素范围在 [-10000, 10000].
 */


func arrayPairSum1(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2{
		sum += nums[i]
	}
	return sum
}

func arrayPairSum(nums []int) int {
	buffer := [20001]int{}

	for _,num := range nums {
		buffer[num+10000]++
	}
	sum := 0
	canAdd := true
	for index,count := range buffer{
		for count > 0 {
			if canAdd {
				sum += index - 10000
			}
			canAdd = !canAdd
			count--
		}
	}
	return sum
}

func TestArrayPairSum(){
	tests := []struct{
		nums []int
		sum int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1,2},
			1,
		},
		{
			[]int{1,3,4,2},
			4,
		},
		{
			[]int{1,3,4,2,6,5},
			9,
		},
		{
			[]int{1,3,4,2,3,3},
			7,
		},
		{
			[]int{1,3,4,2,3,3,4,7},
			11,
		},
	}
	for _,test := range tests{
		fmt.Println(arrayPairSum(test.nums) == test.sum)
	}
}