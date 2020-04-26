package recallTag

import "fmt"

/**
46. 全排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var result [][]int
	var list []int
	_permute(nums, &result, &list)
	return result
}

func _permute(nums []int, result *[][]int, list *[]int) {
	if len(*list) == len(nums) {
		array := make([]int, len(nums))
		copy(array, *list)
		*result = append(*result, array)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !inArray(*list, nums[i]) {
			*list = append(*list, nums[i])
			_permute(nums, result, list)
			*list = (*list)[:len(*list)-1]
		}
	}
}

func inArray(list []int, val int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == val {
			return true
		}
	}
	return false
}

func Testpermute() {
	tests := []struct {
		nums []int
	}{
		{
			[]int{1},
		},
		{
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		result := permute(test.nums)
		fmt.Println(result)
	}
}
