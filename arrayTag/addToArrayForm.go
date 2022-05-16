package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
989. 数组形式的整数加法
对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。

给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。



示例 1：

输入：A = [1,2,0,0], K = 34
输出：[1,2,3,4]
解释：1200 + 34 = 1234
示例 2：

输入：A = [2,7,4], K = 181
输出：[4,5,5]
解释：274 + 181 = 455
示例 3：

输入：A = [2,1,5], K = 806
输出：[1,0,2,1]
解释：215 + 806 = 1021
示例 4：

输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
输出：[1,0,0,0,0,0,0,0,0,0,0]
解释：9999999999 + 1 = 10000000000


提示：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
如果 A.length > 1，那么 A[0] != 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-to-array-form-of-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addToArrayForm(A []int, K int) []int {
	var B []int
	if K < 10 {
		B = []int{K}
	} else {
		count := 5
		B = make([]int, count)
		for K > 0 {
			count--
			B[count] = K % 10
			K /= 10
		}
		B = B[count:]
	}
	aLen, bLen := len(A), len(B)
	maxLen := aLen
	if bLen > maxLen {
		maxLen = bLen
	}
	maxLen++
	C := make([]int, maxLen)
	prev := 0
	for aLen > 0 && bLen > 0 {
		maxLen--
		aLen--
		bLen--
		val := A[aLen] + B[bLen] + prev
		if val >= 10 {
			prev = 1
			C[maxLen] = val - 10
		} else {
			prev = 0
			C[maxLen] = val
		}
	}
	for aLen > 0 {
		maxLen--
		aLen--
		val := A[aLen] + prev
		if val >= 10 {
			prev = 1
			C[maxLen] = val - 10
		} else {
			prev = 0
			C[maxLen] = val
		}
	}
	for bLen > 0 {
		maxLen--
		bLen--
		val := B[bLen] + prev
		if val >= 10 {
			prev = 1
			C[maxLen] = val - 10
		} else {
			prev = 0
			C[maxLen] = val
		}
	}
	if prev == 1 {
		C[0] = 1
		return C
	}
	return C[maxLen:]
}

func TestaddToArrayForm() {
	tests := []struct {
		A []int
		K int
		B []int
	}{
		{
			[]int{0},
			0,
			[]int{0},
		},
		{
			[]int{2},
			3,
			[]int{5},
		},
		{
			[]int{2},
			9,
			[]int{1, 1},
		},
		{
			[]int{1, 2, 0, 0},
			34,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{2, 7, 4},
			181,
			[]int{4, 5, 5},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			1,
			[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, test := range tests {
		B := addToArrayForm(test.A, test.K)
		fmt.Println(common.IntEqual(B, test.B))
	}
}
