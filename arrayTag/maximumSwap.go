package arrayTag

import "fmt"

/**
670. 最大交换
 */

func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}
	buffer := make([]int, 0, 8)
	for num > 0 {
		val := num % 10
		buffer = append(buffer, val)
		num /= 10
	}
	first, second := 0, 0
	i := len(buffer) - 2
	newMaxValue := 0
	//初步筛选出需要被替换的位置
	for ; i >= 0; i-- {
		if buffer[i] > buffer[i+1] {
			first, second = i+1, i
			newMaxValue = buffer[i]
			break
		}
	}
	//找出剩下的最大值
	for ; i >= 0; i-- {
		if buffer[i] >= newMaxValue {
			newMaxValue = buffer[i]
			second = i
		}
	}
	//确保最靠前的元素被替换
	for i := len(buffer) - 1; i >= first; i-- {
		if newMaxValue > buffer[i] {
			first = i
			break
		}
	}
	buffer[first], buffer[second] = buffer[second], buffer[first]
	result := 0
	for i := len(buffer) - 1; i >= 0; i-- {
		result = result*10 + buffer[i]
	}
	return result
}

func TestmaximumSwap() {
	tests := []struct {
		num    int
		result int
	}{
		{
			99901,
			99910,
		},
		{
			999001,
			999100,
		},
		{
			9998001,
			9998100,
		},
		{
			99980019,
			99990018,
		},
		{
			9990001,
			9991000,
		},
		{
			10909091,
			90909011,
		},
		{
			910909091,
			990909011,
		},
		{
			115,
			511,
		},
		{
			1115,
			5111,
		},
		{
			1157,
			7151,
		},
		{
			11577,
			71571,
		},
		{
			11577777,
			71577771,
		},
		{
			1111577777,
			7111577771,
		},
		{
			2736,
			7236,
		},
		{
			12345678,
			82345671,
		},
		{
			987654321,
			987654321,
		},
		{
			987651234,
			987654231,
		},
	}
	for _, test := range tests {
		result := maximumSwap(test.num)
		fmt.Println(result == test.result)
	}
}
