package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
442. 数组中重复的数据
 */

func findDuplicates(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		nums[i] = 0
		for value > 0 {
			if value == nums[value-1] {
				result = append(result, value)
				break
			}
			pos := value - 1
			nums[pos], value = value, nums[pos]
		}
	}
	return result
}

func TestfindDuplicates() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{},
		},
		{
			[]int{2, 2},
			[]int{2},
		},
		{
			[]int{1, 2, 2},
			[]int{2},
		},
		{
			[]int{4, 1, 2, 2},
			[]int{2},
		},
		{
			[]int{1, 1, 2, 2},
			[]int{1, 2},
		},
		{
			[]int{2, 1, 2, 4},
			[]int{2},
		},
		{
			[]int{5, 1, 2, 4, 1},
			[]int{1},
		},
		{
			[]int{5, 1, 2, 2, 1},
			[]int{1, 2},
		},
		{
			[]int{3, 1, 2, 2, 1},
			[]int{2, 1},
		},
		{
			[]int{2, 3, 2, 4, 3},
			[]int{2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{},
		},
		{
			[]int{1, 2, 5, 3, 5},
			[]int{5},
		},
		{
			[]int{1, 3, 2, 4, 6, 2},
			[]int{2},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{3, 2},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1, 9},
			[]int{3, 2},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1, 4},
			[]int{3, 2, 4},
		},
		{
			[]int{2, 3, 7, 7, 8, 2, 4, 1, 3},
			[]int{7, 2, 3},
		},
		{
			[]int{2, 1, 2, 7, 8, 4, 3, 1, 3},
			[]int{2, 1, 3},
		},
		{
			[]int{2, 3, 2, 8, 8, 5, 5, 1, 3},
			[]int{2, 8, 5, 3},
		},
		{
			[]int{2, 4, 2, 8, 7, 8, 3, 1, 3, 1},
			[]int{2, 8, 3, 1},
		},
	}

	for _, test := range tests {
		fmt.Println(common.IntEqualSort(findDuplicates(test.nums), test.result))
	}
}
