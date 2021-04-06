package stringTag

import "fmt"

/**
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	l1, l2 := len(num1), len(num2)
	var next uint8
	buf := make([]byte, l1+1)
	index1 := l1 - 1
	for index2 := l2 - 1; index2 >= 0; index2-- {
		value := (num1[index1] - '0') + (num2[index2] - '0') + next
		if value >= 10 {
			buf[index1+1] = '0' + (value - 10)
			next = 1
		} else {
			buf[index1+1] = '0' + value
			next = 0
		}
		index1--
	}
	for ; index1 >= 0; index1-- {
		value := num1[index1] - '0' + next
		if value >= 10 {
			buf[index1+1] = '0' + (value - 10)
			next = 1
		} else {
			buf[index1+1] = '0' + value
			next = 0
		}
	}
	if next == 1 {
		buf[0] = '1'
		return string(buf)
	}
	return string(buf[1:])
}

func TestaddStrings() {
	tests := []struct {
		num1 string
		num2 string
		ret  string
	}{
		{
			"99999",
			"1",
			"100000",
		},
		{
			"5",
			"5",
			"10",
		},
		{
			"1234",
			"1234",
			"2468",
		},
		{
			"12345",
			"1234",
			"13579",
		},
		{
			"1234",
			"12345",
			"13579",
		},
		{
			"12345",
			"12345",
			"24690",
		},
		{
			"55555",
			"55555",
			"111110",
		},
		{
			"55555",
			"44445",
			"100000",
		},
		{
			"88888",
			"21",
			"88909",
		},
		{
			"88888",
			"22",
			"88910",
		},
	}

	for index, test := range tests {
		ret := addStrings(test.num1, test.num2)
		if ret != test.ret {
			fmt.Println(index, "  ", ret, test.ret)
		}
	}
}
