package dynamicTag

/**
303. 区域和检索 - 数组不可变
 */

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || j > len(this.nums) || i > j {
		return 0
	}
	sum := 0
	for ; i <= j; i++ {
		sum += this.nums[i]
	}
	return sum
}

func TestSumRange() {
	tests := []struct {
		nums []int
		i    int
		j    int
		sum  int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			0,
			2,
			13,
		}, {
			[]int{7, 6, 4, 3, 1},
			1,
			4,
			14,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
			8,
			45,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			3,
			5,
			15,
		},
	}
	for _, test := range tests {
		numRange := Constructor(test.nums)
		fmt.Println(numRange.SumRange(test.i, test.j) == test.sum)
	}
}
