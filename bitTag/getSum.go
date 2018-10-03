package bitTag

/**
371. 两整数之和
分析：a^b忽略都存在1的二进制，不同的二进制直接合并，a&b有相同二进制的直接进1，直到没有相同的二进制位位置
 */

func getSum(a int, b int) int {
	if a & b == 0 {
		return a|b
	}
	return getSum(a^b, (a&b)<<1)
}