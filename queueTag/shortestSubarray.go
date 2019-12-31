package queueTag

/**
862. 和至少为 K 的最短子数组
返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。

如果没有和至少为 K 的非空子数组，返回 -1 。


示例 1：

输入：A = [1], K = 1
输出：1
示例 2：

输入：A = [1,2], K = 4
输出：-1
示例 3：

输入：A = [2,-1,2], K = 3
输出：3
 

提示：

1 <= A.length <= 50000
-10 ^ 5 <= A[i] <= 10 ^ 5
1 <= K <= 10 ^ 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func shortestSubarray(A []int, K int) int {
	sum := A[0]
	if sum >= K {
		return 1
	}
	index := 0
	min := len(A) + 1
	for i := 1; i < len(A); i++ {
		sum += A[i]
		if sum >= K {
			for index < i {
				if i-index+1 < min {
					min = i - index + 1
				}
				sum -= A[index]
				index++
				if sum < K {
					break
				}
			}
		}
	}
	if min > len(A) {
		return -1
	}
	return min
}
