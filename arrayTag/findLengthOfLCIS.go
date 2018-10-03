package arrayTag

import "fmt"

/**
674. 最长连续递增序列
 */

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen, l := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			l++
			if l > maxLen {
				maxLen = l
			}
		} else {
			l = 1
		}
	}
	return maxLen
}

func TestfindLengthOfLCIS() {
	tests := []struct {
		nums  []int
		count int
	}{
		{
			[]int{1, 1, 1, 1, 1},
			1,
		},
		{
			[]int{1, 3, 5, 4, 7},
			3,
		},
		{
			[]int{1, 3, 5, 4, 7, 8, 9},
			4,
		},
		{
			[]int{1, 3, 5, 4, 7, 8, 9, 1, 2, 3},
			4,
		},
	}

	for _, test := range tests {
		count := findLengthOfLCIS(test.nums)
		fmt.Println(count == test.count)
	}
}
