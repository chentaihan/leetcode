package bitTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
260. 只出现一次的数字 III
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
分析：异或运算得到目标两数的异或结果，这个数的某一位只属于两个数中其中的一个，通过这个位将数组分成两份，
这个两个数就分到两个数组中去了，分别异或运算就能得到目标两个数
*/

func singleNumber2(nums []int) []int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum ^= nums[i]
	}
	flag := 1
	for sum&flag == 0 {
		flag <<= 1
	}
	ret := make([]int, 2)
	ret[0], ret[1] = sum, sum
	for _, num := range nums {
		if num&flag > 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}
	return ret
}

func TestsingleNumber2() {
	tests := []struct {
		nums []int
		ret  []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 1, 4, 5, 6},
			[]int{2, 3},
		},
		{
			[]int{1, 2, 2, 4, 5, 6, 3, 4, 5, 6},
			[]int{1, 3},
		},
	}
	for _, test := range tests {
		ret := singleNumber2(test.nums)
		fmt.Println(common.IntEqual(ret, test.ret))
	}
}
