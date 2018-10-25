package bitTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
338. 比特位计数
 */

//通过val = val & (val - 1)计算二进制位中1的个数
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		count := 0
		val := i
		for ; val > 0; val = val & (val - 1) {
			count++
		}
		result[i] = count
	}
	return result
}

//规律：ret[i] = ret[i/2] + i%2，如果刚好是2倍，则二进制为1的位数是一样的，多出来1就加一，即i%2
func countBitsEx(num int) []int {
	ret := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ret[i] = ret[i/2] + i%2
	}
	return ret
}

func TestcountBits() {
	tests := []struct {
		num    int
		result []int
	}{
		{
			0,
			[]int{0},
		},
		{
			1,
			[]int{0, 1},
		},
		{
			2,
			[]int{0, 1, 1},
		},
		{
			3,
			[]int{0, 1, 1, 2},
		},
		{
			4,
			[]int{0, 1, 1, 2, 1},
		},
		{
			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
		{
			6,
			[]int{0, 1, 1, 2, 1, 2, 2},
		},
		{
			7,
			[]int{0, 1, 1, 2, 1, 2, 2, 3},
		},
		{
			8,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
		{
			9,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2},
		},
		{
			10,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
		{
			11,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3},
		},
	}
	for _, test := range tests {
		result := countBits(test.num)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
