package stringTag

/**
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 */

func multiply(num1 string, num2 string) string {
	buffer := []byte{}
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {

		}
	}
	return string(buffer)
}

func charAdd(c1, c2 byte) (byte, byte) {
	val := c1 - '0' + c2 - '0'
	if val < 10 {
		return '0', '0' + val
	}
	return '1', '0' + (val - 10)
}

func charMult(c1, c2 byte) (byte, byte) {
	val := (c1 - '0') * (c2 - '0')
	if val < 10 {
		return '0', '0' + val
	}
	return '0' + (val / 10), '0' + (val % 10)
}

func stringAdd(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	minLen, maxlen := l1, l1
	if l2 < minLen {
		minLen = l2
		maxlen = l1 + 1
	} else {
		minLen = l1
		maxlen = l2 + 1
	}
	buffer := make([]byte, maxlen)
	leftC := byte('0')
	for ; minLen >= 0; minLen-- {
		decade, theunit := charAdd(s1[minLen], s2[minLen])
		if leftC != '0' {
			buffer[minLen] = theunit
		} else {
			decade2, theunit := charAdd(theunit, leftC)
			_, leftC = charAdd(decade, decade2)
			buffer[minLen] = theunit
		}
	}
	return string(buffer)
}
