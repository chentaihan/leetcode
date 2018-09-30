package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
229. 求众数 II
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
 */

func majorityElement2(nums []int) []int {
	num1, num2 := 0, 0
	count1, count2 := 0, 0
	for _, num := range nums {
		if count1 == 0 && num != num2 {
			num1 = num
			count1++
			continue
		}
		if count2 == 0 && num != num1 {
			num2 = num
			count2++
			continue
		}
		if num == num1 {
			count1++
		}else if num == num2 {
			count2++
		}else{
			count1--
			count2--
		}
	}
	result := []int{}
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		}
	}
	if float64(count1) > float64(len(nums))/3.0 {
		result = append(result, num1)
	}
	if float64(count2) > float64(len(nums))/3.0 {
		result = append(result, num2)
	}
	return result
}

func TestmajorityElement2() {
	tests := []struct {
		nums   []int
		result []int
	}{
		//{
		//	[]int{3, 3, 4},
		//	[]int{3},
		//},
		{
			[]int{1, 2, 2, 3, 2, 1, 1, 3},
			[]int{2, 1},
		},
		{
			[]int{3, 3, 3, 4, 2, 2, 2},
			[]int{3, 2},
		},
		{
			[]int{3, 3, 3, 4, 2, 2, 4, 2},
			[]int{3, 2},
		},
	}
	for _, test := range tests {
		fmt.Println(common.IntEqual(majorityElement2(test.nums), test.result))
	}
}
