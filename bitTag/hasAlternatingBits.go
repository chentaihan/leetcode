package bitTag

import "fmt"

/**
693. 交替位二进制数
 */

func hasAlternatingBits(n int) bool {
	preFlag := n & 1
	for n > 0 {
		n >>= 1
		flag := n & 1
		if preFlag == flag {
			return false
		}
		preFlag = flag
	}
	return true
}

func TesthasAlternatingBits() {
	tests := []struct{
		n int
		ret bool
	}{
		{
			0,true,
		},
		{
			1,true,
		},
		{
			2,true,
		},
		{
			3,false,
		},
		{
			4,false,
		},
		{
			5,true,
		},
		{
			6,false,
		},
		{
			7,false,
		},
		{
			8,false,
		},
		{
			10,true,
		},
		{
			21,true,
		},
		{
			22,false,
		},
		{
			42,true,
		},
		{
			49,false,
		},

	}
	for _,test := range tests {
		ret := hasAlternatingBits(test.n)
		fmt.Println(ret == test.ret)
	}
}
