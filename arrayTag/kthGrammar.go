package arrayTag

import "fmt"

/**
779. 第K个语法符号
在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。

给定行数 N 和序数 K，返回第 N 行中第 K个字符。（K从1开始）


例子:

输入: N = 1, K = 1
输出: 0

输入: N = 2, K = 1
输出: 0

输入: N = 2, K = 2
输出: 1

输入: N = 4, K = 5
输出: 1

解释:
第一行: 0
第二行: 01
第三行: 0110
第四行: 01101001
第五行: 01 10 10 01 10 01 01 10


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/k-th-symbol-in-grammar
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func kthGrammar(N int, K int) int {
	K -= 1
	K ^= K >> 1
	K ^= K >> 2
	K = (K & 0x11111111) * 0x11111111
	return (K >> 28) & 1
}

func TestkthGrammar() {
	tests := []struct {
		N int
		K int
		V int
	}{
		{
			1, 1, 0,
		},
		{
			2, 1, 0,
		},
		{
			2, 2, 1,
		},
		{
			4, 5, 1,
		},
		{
			4, 6, 0,
		},
		{
			5, 3, 1,
		},
		{
			5, 4, 0,
		},
		{
			5, 5, 1,
		},
		{
			5, 6, 0,
		},
	}
	for _, test := range tests {
		v := kthGrammar(test.N, test.K)
		if v != test.V {
			fmt.Println("error ", v, test.V)
		}
	}
}
