package arrayTag

import "fmt"

/**
915. 分割数组
给定一个数组 A，将其划分为两个不相交（没有公共元素）的连续子数组 left 和 right， 使得：

left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 要尽可能小。
在完成这样的分组后返回 left 的长度。可以保证存在这样的划分方法。

 

示例 1：

输入：[5,0,3,8,6]
输出：3
解释：left = [5,0,3]，right = [8,6]
示例 2：

输入：[1,1,1,0,6,12]
输出：4
解释：left = [1,1,1,0]，right = [6,12]
 

提示：

2 <= A.length <= 30000
0 <= A[i] <= 10^6
可以保证至少有一种方法能够按题目所描述的那样对 A 进行划分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func partitionDisjoint(A []int) int {
	minValue, leftMax, max := A[0], A[0],0
	for i := 1; i < len(A); i++ {
		if A[i] > leftMax {
			leftMax = A[i]
		}
		if A[i] < minValue {
			minValue = leftMax
			max = i
		}
	}
	return max + 1
}

func TestpartitionDisjoint() {
	tests := []struct {
		A     []int
		count int
	}{
		{
			[]int{5, 0, 3, 8, 6},
			3,
		},
		{
			[]int{1, 1, 1, 0, 6, 12},
			4,
		},
		{
			[]int{1, 1, 1, 2, 6, 12},
			1,
		},
		{
			[]int{0, 1, 1, 1, 2, 6, 12},
			1,
		},
		{
			[]int{1, 1, 1, 2, 6, 0, 12},
			6,
		},
		{
			[]int{29, 33, 6, 4, 42, 0, 10, 22, 62, 16, 46, 75, 100, 67, 70, 74, 87, 69, 73, 88},
			11,
		},
	}
	for _, test := range tests {
		count := partitionDisjoint(test.A)
		fmt.Println(count == test.count)
	}
}
