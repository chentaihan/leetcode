package arrayTag

import "fmt"

/**
941. 有效的山脉数组
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

A.length >= 3
在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[B.length - 1]
 

示例 1：

输入：[2,1]
输出：false
示例 2：

输入：[3,5,5]
输出：false
示例 3：

输入：[0,3,2,1]
输出：true
 

提示：

0 <= A.length <= 10000
0 <= A[i] <= 10000 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-mountain-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	isTop := false
	maxValue := 0
	for i := 1; i < len(A); i++ {
		if !isTop {
			if A[i] <= A[i-1] {
				isTop = true
				maxValue = A[i-1]
			}
		} else {
			if A[i] >= A[i-1] {
				return false
			}
		}
	}
	return isTop && A[len(A)-1] < maxValue && A[0] < maxValue
}

func TestvalidMountainArray() {
	tests := []struct {
		A []int
		r bool
	}{
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			false,
		},
		{
			[]int{1, 2, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5, 5},
			false,
		},
	}
	for _, test := range tests {
		r := validMountainArray(test.A)
		fmt.Println(r == test.r)
	}
}
