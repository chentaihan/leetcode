package bitTag

import (
	"fmt"
)

/**
405. 数字转换为十六进制数
 */

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = 1<<32 + num
	}

	ret := ""
	for num > 0 {
		val := num % 16
		if val >= 0 && val <= 9 {
			ret = string('0'+val) + ret
		} else {
			ret = string('a'+(val-10)) + ret
		}
		num = (num - val) / 16
	}
	return ret
}

func TesttoHex() {
	tests := []struct {
		num int
		ret string
	}{
		{
			-1,
			"ffffffff",
		},
		{
			-2,
			"fffffffe",
		},
		{
			-16,
			"ffffffef",
		},
		{
			0,
			"0",
		},
		{
			15,
			"f",
		},
		{
			16,
			"10",
		},
		{
			17,
			"11",
		},
		{
			19,
			"13",
		},
		{
			26,
			"1a",
		},
		{
			31,
			"1f",
		},
		{
			32,
			"20",
		},
		{
			34,
			"22",
		},
		{
			160,
			"a0",
		},
		{
			176,
			"b0",
		}, {
			170,
			"aa",
		},
		{
			256,
			"100",
		},
		{
			273,
			"111",
		},
		{
			693,
			"2b5",
		},
	}
	for _, test := range tests {
		fmt.Println(toHex(test.num) == test.ret)
	}
}
