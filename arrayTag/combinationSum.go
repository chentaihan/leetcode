package arrayTag

import (
	"fmt"
	"sort"
)

/**
39. 组合总和
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
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	return combineSum(candidates, target)
}

func combineSum(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 || target < candidates[0] {
		return result
	}
	for i, num := range candidates {
		if target == num {
			item := []int{num}
			result = append(result, item)
		} else if num > target {
			break
		} else {
			res := combineSum(candidates[i:], target-num)
			for _, list := range res {
				item := append([]int{num}, list...)
				result = append(result, item)
			}
		}
	}
	return result
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
