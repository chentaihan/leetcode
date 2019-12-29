package sortTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
57. 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:

输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insert-interval
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	for index := 0; index < len(intervals); index++ {
		flag := cross(intervals[index], newInterval)
		switch flag {
		case 0: //left
			res = append(res, newInterval)
			return append(res, intervals[index:]...)
		case 1: //central
			if intervals[index][0] < newInterval[0] {
				newInterval[0] = intervals[index][0]
			}
			if intervals[index][1] > newInterval[1] {
				newInterval[1] = intervals[index][1]
			}
		case 2: //right
			res = append(res, intervals[index])
		}
	}
	return append(res, newInterval)
}

func cross(arr1, arr2 []int) int {
	if arr2[1] < arr1[0] { //left
		return 0
	}
	if arr2[0] > arr1[1] { //right
		return 2
	}
	return 1 //central
}

func TestInsert() {
	tests := []struct {
		intervals [][]int
		newInts   []int
		res       [][]int
	}{
		{
			[][]int{},
			[]int{0, 1},
			[][]int{{0, 1}},
		},
		{
			[][]int{{2, 3}, {5, 9}},
			[]int{0, 1},
			[][]int{{0, 1}, {2, 3}, {5, 9}},
		},
		{
			[][]int{{2, 3}, {5, 9}},
			[]int{0, 13},
			[][]int{{0, 13}},
		},
		{
			[][]int{{2, 3}, {5, 9}, {20, 22}},
			[]int{0, 13},
			[][]int{{0, 13}, {20, 22}},
		},
		{
			[][]int{{2, 3}, {5, 9}, {20, 22}},
			[]int{0, 43},
			[][]int{{0, 43}},
		},
		{
			[][]int{{2, 3}, {5, 9}, {20, 22}},
			[]int{3, 21},
			[][]int{{2, 22}},
		},
		{
			[][]int{{2, 3}, {5, 9}, {20, 22}},
			[]int{2, 22},
			[][]int{{2, 22}},
		},
		{
			[][]int{{2, 3}, {5, 9}},
			[]int{0, 2},
			[][]int{{0, 3}, {5, 9}},
		},
		{
			[][]int{{2, 3}, {5, 9}},
			[]int{0, 4},
			[][]int{{0, 4}, {5, 9}},
		},
		{
			[][]int{{2, 4}, {8, 12}},
			[]int{3, 6},
			[][]int{{2, 6}, {8, 12}},
		},
		{
			[][]int{{2, 4}, {8, 12}},
			[]int{3, 9},
			[][]int{{2, 12}},
		},
		{
			[][]int{{2, 4}, {8, 12}},
			[]int{13, 19},
			[][]int{{2, 4}, {8, 12}, {13, 19}},
		},
	}

	for _, test := range tests {
		res := insert(test.intervals, test.newInts)
		fmt.Println(common.IntEqualEx(res, test.res))
	}
}
