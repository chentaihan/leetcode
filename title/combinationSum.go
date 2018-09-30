package title

import (
	"sort"
	"fmt"
)

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
https://leetcode-cn.com/problems/combination-sum/description/
 */

//二分查找
func binarySearch(nums []int, target int) int {
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

	return -1
}

func copyArray(nums []int) []int {
	array := make([]int, len(nums))
	copy(array, nums)
	return array
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := make([][]int, 0)
	combineSum(candidates, []int{target}, &result)
	return result
}

func combineSum(candidates []int, target []int, result *[][]int) {
	for _, num := range candidates {
		index := binarySearch(candidates, target[len(target)-1])
		if index >= 0 {
			*result = append(*result, target)
		}

		newValue := target[len(target)-1] - num
		if newValue >= candidates[0] {
			target = copyArray(target)
			target[len(target)-1] = num
			target = append(target, newValue)
			combineSum(candidates, target, result)
		}
	}
}

func TestCombinationSum() {
	tests := []struct {
		candidates []int
		target     int
		result     [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{{7}, {2, 2, 3}},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}
	for _, test := range tests {
		result := combinationSum(test.candidates, test.target)
		for _, nums := range result {
			for _, num := range nums {
				fmt.Print(num, " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
