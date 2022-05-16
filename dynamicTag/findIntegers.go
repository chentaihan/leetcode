package dynamicTag

/*
600. 不含连续1的非负整数

给定一个正整数 n，找出小于或等于 n 的非负整数中，其二进制表示不包含 连续的1 的个数。

示例 1:

输入: 5
输出: 5
解释:
下面是带有相应二进制表示的非负整数<= 5：
0 : 0
1 : 1
2 : 10
3 : 11
4 : 100
5 : 101
其中，只有整数3违反规则（有两个连续的1），其他5个满足规则。
说明: 1 <= n <= 10 << 9
*/

func findIntegers(num int) int {
	newNum := num
	var buffer []int
	for newNum > 0 {
		buffer = append(buffer, newNum&1)
		newNum >>= 1
	}
	l := len(buffer)
	count := 0
	for i := 0; i < l; i++ {
		if buffer[i] == 1 {
			buffer[i] = 0
		} else {
			buffer[i] = 1
			for j := i + 1; j < l; j++ {
				if buffer[i] == 1 {
					buffer[j] = 0
					break
				} else {
					buffer[j] = 1
				}
			}
		}

		isok := true
		for j := i + 1; j < l; j++ {
			if buffer[i] == 0 {
				continue
			}
			if i+1 < l && buffer[i+1] == 1 {
				isok = false
				break
			}
		}
		if isok {
			count++
		}
	}

	return 0
}
