package arrayTag

import (
	"fmt"
	"sort"
)

/**
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	return combineSum2(candidates, target)
}

func combineSum2(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 || target < candidates[0] {
		return result
	}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		j := i
		for ; j < len(candidates)-1; j++ {
			if candidates[j] != candidates[j+1] {
				break
			}
		}
		var left []int
		sum := 0
		for k := i; k <= j; k++ {
			left = append(left, candidates[k])
			sum += candidates[k]
			if sum == target {
				result = append(result, left)
				break
			} else {
				res := combineSum2(candidates[j+1:], target-sum)
				for _, list := range res {
					item := append(left, list...)
					result = append(result, item)
				}
			}
		}
		i = j
	}
	return result
}

func TestCombinationSum2() {
	tests := []struct {
		candidates []int
		target     int
		result     [][]int
	}{
		{
			[]int{1, 1},
			1,
			[][]int{{1}},
		},
		{
			[]int{1, 2, 2, 5},
			5,
			[][]int{{1, 2, 2}, {5}},
		},
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{{7}},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{{3, 5}},
		},
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			[]int{1, 1, 2, 5, 6, 7, 10},
			8,
			[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
	}
	for _, test := range tests {
		result := combinationSum2(test.candidates, test.target)

		fmt.Println(result)
	}
}
