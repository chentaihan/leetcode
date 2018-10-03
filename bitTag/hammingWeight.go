package bitTag

/**
191. 位1的个数
 */
func hammingWeight(n uint32) int {
	count := 0
	for ; n > 0; n = n & (n - 1) {
		count++
	}
	return count
}
