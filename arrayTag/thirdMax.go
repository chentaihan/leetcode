package arrayTag

import (
	"fmt"
)

/**
414. 第三大的数
 */

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	ret, count, i := _thirdMax(nums)
	if count < 3 {
		return ret[count-1]
	}
	for ; i < len(nums); i++ {
		if nums[i] > ret[0] {
			index := 1
			for ; index < 3; index++ {
				if nums[i] < ret[index] {
					break
				}
				if nums[i] == ret[index] {
					goto end
				}
			}
			for j := 1; j < index; j++ {
				ret[j-1] = ret[j]
			}
			ret[index-1] = nums[i]
		}
	end:
	}
	return ret[0]
}

//填充3个数
func _thirdMax(nums []int) ([]int, int, int) {
	ret := make([]int, 3)
	ret[0] = nums[0]
	index := 1
	i := 1
	for ; i < len(nums) && index < 3; i++ {
		j := 0
		for ; j < index; j++ {
			if nums[i] < ret[j] {
				break
			}
			if nums[i] == ret[j] {
				goto end
			}
		}
		for k := index; k > j; k-- {
			ret[k] = ret[k-1]
		}
		ret[j] = nums[i]
		index++
	end:
	}

	return ret, index, i
}

func TestthirdMax() {
	tests := []struct {
		nums []int
		ret  int
	}{
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{2, 1},
			2,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 2, 3},
			1,
		},
		{
			[]int{1, 2, 2},
			2,
		},
		{
			[]int{3, 1, 2, 3},
			1,
		},
		{
			[]int{3, 2, 1, 2, 3},
			1,
		},
		{
			[]int{1, 2, 2, 5, 3, 5},
			2,
		},
		{
			[]int{4, 3, 2, 1, 2, 3},
			2,
		},
		{
			[]int{5, 1, 2, 3, 4, 5, 3},
			3,
		},
		{
			[]int{5, 1, 3, 2, 4, 6, 3},
			4,
		},
		{
			[]int{5, 1, 3, 7, 2, 4, 6, 3},
			5,
		},
		{
			[]int{8, 5, 1, 3, 7, 2, 4, 6, 3},
			6,
		},
		{
			[]int{8, 5, 1, 3, 7, 2, 7, 4, 6, 3},
			6,
		},
		{
			[]int{8, 5, 1, 3, 7, 2, 7, 4, 6, 3, 8},
			6,
		},
		{
			[]int{8, 5, 1, 3, 7, 2, 7, 4, 6, 3, 8, 9},
			7,
		},
		{
			[]int{8, 5, 1, 3, 7, 2, 7, 4, 6, 3, 8, 9, 8},
			7,
		},
		{
			[]int{10, 8, 5, 1, 3, 7, 2, 7, 4, 6, 3, 8, 9, 8},
			8,
		},
		{
			[]int{10, 8, 5, 1, 3, 7, 2, 7, 4, 6, 3, 8, 9, 8, 10},
			8,
		},
	}
	for _, test := range tests {
		fmt.Println(thirdMax(test.nums) == test.ret)
	}
}
