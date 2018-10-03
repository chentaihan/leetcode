package bitTag

import "fmt"

/**
190. 颠倒二进制位
 */

func reverseBits(n uint32) uint32 {
	result := uint32(0)
	for i := uint32(1); i <= 32; i++ {
 		result += (n&uint32(1)) << (uint32(32) - i)
 		n >>=1
	}
	return result
}


func TestreverseBits(){
	tests := []struct{
		n uint32
		r uint32
	}{
		{
			43261596,
			964176192,
		},
	}
	for _,test := range tests{
		result := reverseBits(test.n)
		fmt.Println(result == test.r)
	}
}