package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
448. 找到所有数组中消失的数字
 */

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			continue
		}
		value := nums[i]
		nums[i] = 0
		for value > 0 {
			if value == nums[value-1] {
				break
			}
			pos := value - 1
			nums[pos], value = value, nums[pos]
		}
	}
	result := []int{}
	for index, val := range nums {
		if val == 0 {
			result = append(result, index+1)
		}
	}
	return result
}

func TestfindDisappearedNumbers() {
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
			[]int{1, 2},
			[]int{},
		},
		{
			[]int{1, 2, 2},
			[]int{3},
		},
		{
			[]int{4, 1, 2, 2},
			[]int{3},
		},
		{
			[]int{2, 1, 2, 2},
			[]int{3, 4},
		},
		{
			[]int{2, 1, 2, 4},
			[]int{3},
		},
		{
			[]int{5, 1, 2, 4, 1},
			[]int{3},
		},
		{
			[]int{5, 1, 2, 2, 1},
			[]int{3, 4},
		},
		{
			[]int{2, 1, 2, 2, 1},
			[]int{3, 4, 5},
		},
		{
			[]int{2, 2, 2, 2, 2},
			[]int{1, 3, 4, 5},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{2, 3, 4, 5},
		},
		{
			[]int{5, 5, 5, 5, 5},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 3, 2, 4, 6, 2},
			[]int{5},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1, 9},
			[]int{5, 6},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1, 3},
			[]int{5, 6, 9},
		},
		{
			[]int{2, 3, 2, 7, 8, 2, 3, 1, 3},
			[]int{4, 5, 6, 9},
		},
		{
			[]int{2, 3, 2, 7, 8, 2, 3, 3, 3},
			[]int{1, 4, 5, 6, 9},
		},
		{
			[]int{2, 3, 2, 8, 8, 2, 3, 3, 3},
			[]int{1, 4, 5, 6, 7, 9},
		},
		{
			[]int{2, 3, 2, 8, 8, 2, 3, 3, 3, 1},
			[]int{4, 5, 6, 7, 9, 10},
		},
	}

	for _, test := range tests {
		fmt.Println(common.IntEqual(findDisappearedNumbers(test.nums), test.result))
	}
}
