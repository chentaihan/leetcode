package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
283. 移动零
 */

func moveZeroes(nums []int) {
	l := len(nums)
	for index := 0; index < l; index++ {
		if nums[index] == 0 {
			j := index + 1
			for ; j < l; j++ {
				if nums[j] != 0 {
					nums[index], nums[j] = nums[j], nums[index]
					break
				}
			}
			if j == l {
				break
			}
		}
	}
}

func TestmoveZeroes() {
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
			[]int{1},
		},
		{
			[]int{0, 2},
			[]int{2, 0},
		},
		{
			[]int{0, 1, 2, 2},
			[]int{1, 2, 2, 0},
		},
		{
			[]int{4, 0, 1, 2, 2},
			[]int{4, 1, 2, 2, 0},
		},
		{
			[]int{1, 1, 2, 2, 0},
			[]int{1, 1, 2, 2, 0},
		},
		{
			[]int{2, 1, 2, 0, 4},
			[]int{2, 1, 2, 4, 0},
		},
		{
			[]int{5, 7, 0, 2, 4, 0, 1},
			[]int{5, 7, 2, 4, 1, 0, 0},
		},
		{
			[]int{0, 5, 1, 7, 12, 0, 1},
			[]int{5, 1, 7, 12, 1, 0, 0},
		},
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{2, 0, 3, 0, 12, 0, 7, 0, 8, 0, 15, 0, 4, 0, 11, 0, 1, 0},
			[]int{2, 3, 12, 7, 8, 15, 4, 11, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.nums)
		fmt.Println(common.IntEqual(test.nums, test.result))
	}
}
