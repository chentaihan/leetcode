package bitTag

/**
231. 2的幂
分析：2的幂和其前一个数二进制位没有交集
 */

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}