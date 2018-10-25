package bitTag

import "fmt"

/**
476. 数字的补数
 */

func findComplement(num int) int {
	count := uint(0)
	val := num
	for ; val > 0; val >>= 1 {
		count++
	}
	return (1<<count) - 1 - num
}

func TestfindComplement() {
	tests := []struct {
		num int
		ret int
	}{
		{
			1,
			0,
		},
		{
			2,
			1,
		},
		{
			3,
			0,
		},
		{
			4,
			3,
		},
		{
			5,
			2,
		},
		{
			15,
			0,
		},
		{
			16,
			15,
		},
		{
			32,
			31,
		},

	}
	for _,test := range tests{
		fmt.Println(findComplement(test.num) == test.ret)
	}
}
