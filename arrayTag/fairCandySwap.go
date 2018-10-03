package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
888. 公平的糖果交换
 */

func fairCandySwap(A []int, B []int) []int {
	aSum, bSum := 0, 0
	for _, num := range A {
		aSum += num
	}
	m := make(map[int]struct{})
	for _, num := range B {
		bSum += num
		m[num] = struct{}{}
	}
	addValue := (aSum - bSum) / 2
	for _, num := range A {
		if _, exist := m[num-addValue]; exist {
			return []int{num, num - addValue}
		}
	}
	return []int{}
}

func TestfairCandySwap() {
	tests := []struct {
		A []int
		B []int
		R []int
	}{
		{
			[]int{1, 1},
			[]int{2, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2},
			[]int{2, 3},
			[]int{1, 2},
		},
		{
			[]int{2},
			[]int{1, 3},
			[]int{2, 3},
		},
	}
	for _, test := range tests {
		result := fairCandySwap(test.A, test.B)
		fmt.Println(common.IntEqual(result, test.R))
	}
}
