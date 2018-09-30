package arrayTag

import "fmt"

/**
169. 求众数
*/

func majorityElement(nums []int) int {
	result := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if result == nums[i] {
			count++
			if count > len(nums)/2 {
				return result
			}
		} else {
			if count == 1 {
				result = nums[i]
			} else {
				count--
			}
		}
	}
	return result
}

func TestmajorityElement() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{3, 3, 4},
			3,
		},
		{
			[]int{1, 2, 2, 2, 3, 4, 2},
			2,
		},
		{
			[]int{1, 3, 2, 3, 3, 4, 3},
			3,
		},
		{
			[]int{1, 3, 2, 3, 4, 3, 5, 3, 3},
			3,
		},
		{
			[]int{3, 1, 3, 2, 3, 4, 3, 5, 3},
			3,
		},
	}
	for _, test := range tests {
		fmt.Println(majorityElement(test.nums) == test.result)
	}
}
