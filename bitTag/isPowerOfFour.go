package bitTag

import "fmt"

/**
342. 4的幂
解析：2863311531&num == 0二进制位为1的都是偶数为，num&(num-1)且只有一位是1
 */

func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	return 2863311531&num == 0 && num&(num-1) == 0
}

func TestisPowerOfFour() {
	tests := []struct {
		num    int
		result bool
	}{
		{
			0,
			false,
		},
		{
			1,
			true,
		},
		{
			4,
			true,
		},
		{
			5,
			false,
		},
		{
			16,
			true,
		},
		{
			17,
			false,
		},
		{
			20,
			false,
		},
		{
			64,
			true,
		},
		{
			128,
			false,
		},
		{
			256,
			true,
		},
		{
			512,
			false,
		},
		{
			1024,
			true,
		},
		{
			2048,
			false,
		}, {
			4096,
			true,
		},
	}
	for _, test := range tests {
		result := isPowerOfFour(test.num)
		fmt.Println(result == test.result)
	}
}
