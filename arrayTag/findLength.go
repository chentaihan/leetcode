package arrayTag

import "fmt"

/**
718. 最长重复子数组
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例 1:

输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
说明:

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：和最长公共子字符串一样，动态规划，同一个数组保存已经保存的字符串数量，由于使用一维数组，为了避免覆盖情况，得从后往前比较
*/

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	buf := make([]int, len(B))
	max := 0
	for i := 0; i < len(A); i++ {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				if j > 0 && buf[j-1] > 0 {
					buf[j] = buf[j-1] + 1
				} else {
					buf[j] = 1
				}
				if max < buf[j] {
					max = buf[j]
				}
			} else {
				buf[j] = 0
			}
		}
	}
	return max
}

func TestfindLength() {
	tests := []struct {
		A []int
		B []int
		l int
	}{
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1},
			[]int{2},
			0,
		},
		{
			[]int{1},
			[]int{1},
			1,
		},
		{
			[]int{1, 3, 4, 5, 7},
			[]int{2, 3, 4, 5, 8},
			3,
		},
		{
			[]int{1, 3, 4, 5, 7},
			[]int{1, 3, 4, 5, 8},
			4,
		},
		{
			[]int{1, 3, 4, 5, 7},
			[]int{1, 3, 4, 5, 7},
			5,
		},
		{
			[]int{1, 3, 4, 5, 7},
			[]int{1, 2, 4, 5, 8},
			2,
		},
		{
			[]int{1, 3, 4, 5, 7},
			[]int{1, 2, 4, 6, 8},
			1,
		},
	}
	for _, test := range tests {
		l := findLength(test.A, test.B)
		fmt.Println(l == test.l)
	}
}
