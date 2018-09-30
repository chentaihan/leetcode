package stackTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
496. 下一个更大元素 I
 */

func nextGreaterElement(findNums []int, nums []int) []int {
	if len(findNums) == 0 || len(nums) == 0 {
		return []int{}
	}
	m := make(map[int]int, len(nums))
	for key, val := range nums {
		m[val] = key
	}
	buffer := make([]int, len(findNums))
	for key, val := range findNums {
		nextVal := -1
		for start := m[val]; start < len(nums); start++ {
			if nums[start] > val {
				nextVal = nums[start]
				break
			}
		}
		buffer[key] = nextVal
	}
	return buffer
}

func TestnextGreaterElement() {
	tests := []struct {
		findNums []int
		nums     []int
		result   []int
	}{
		{
			[]int{7, 1, 5},
			[]int{7, 1, 5, 3, 6, 4},
			[]int{-1, 5, 6},
		},
		{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	}
	for _, test := range tests {
		fmt.Println(common.IntEqual(nextGreaterElement(test.findNums, test.nums), test.result))
	}
}
