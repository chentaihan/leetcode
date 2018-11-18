package arrayTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
167. 两数之和 II - 输入有序数组
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

分析：通过target/2将数组分成两半，这两个数肯定分别出现在这两个数组中，接壤的地方做特殊处理
 */

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}
	//numbers[avgIndex]出现在两个数组中，需要做特殊处理，考虑相同数的情况
	avgIndex := binarySearchEx(numbers, target/2)
	for i := 0; i <= avgIndex; i++ {
		val := target - numbers[i]
		index := binarySearch(numbers[avgIndex:], val)
		if index >= 0 {
			if i != avgIndex+index {
				return []int{i + 1, avgIndex + index + 1}
			}
			if i+1 < len(numbers) && numbers[i] == numbers[i+1] {
				return []int{i + 1, avgIndex + index + 2}
			}
		}
	}
	return []int{}
}

//二分查找
func binarySearchEx(nums []int, target int) int {
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return middle
}

func TesttwoSum() {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{

		{
			[]int{1, 2, 3, 4, 4, 9, 56, 90},
			8,
			[]int{4, 5},
		},
		{
			[]int{1, 2, 3, 4, 4, 9, 56, 90, 90},
			180,
			[]int{8, 9},
		},
		{
			[]int{1, 2, 2, 3, 4, 4, 9, 56, 90, 90},
			180,
			[]int{9, 10},
		},
		{
			[]int{1, 2},
			3,
			[]int{1, 2},
		},
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			[]int{2, 7, 11, 15},
			17,
			[]int{1, 4},
		},
		{
			[]int{2, 7, 11, 15},
			17,
			[]int{1, 4},
		},
		{
			[]int{2, 7, 11, 15},
			18,
			[]int{2, 3},
		},
		{
			[]int{2, 7, 11, 15},
			1,
			[]int{},
		},
		{
			[]int{2, 7, 11, 15, 19, 21, 35},
			34,
			[]int{4, 5},
		},
		{
			[]int{2, 7, 11, 15, 90, 100},
			102,
			[]int{1, 6},
		},
		{
			[]int{2, 7, 11, 15, 90, 100, 300},
			115,
			[]int{4, 6},
		},
		{
			[]int{2, 7, 11, 15, 90, 100, 300},
			1,
			[]int{},
		},
		{
			[]int{2, 7, 11, 15, 90, 100, 300},
			400,
			[]int{6, 7},
		},
		{
			[]int{2, 7, 11, 15, 90, 100, 300, 500},
			800,
			[]int{7, 8},
		},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
