package sortTag

/**
56. 合并区间
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"sort"
)

type IntsSlice [][]int

func (p IntsSlice) Len() int { return len(p) }
func (p IntsSlice) Less(i, j int) bool {
	return p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}
func (p IntsSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(IntsSlice(intervals))
	var ret [][]int
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= cur[1] {
			if intervals[i][1] > cur[1] {
				cur[1] = intervals[i][1]
			}
		} else {
			ret = append(ret, cur)
			cur = intervals[i]
		}
	}
	if len(ret) == 0 || cur[0] > ret[len(ret)-1][1] {
		ret = append(ret, cur)
	}
	return ret
}

func TestMerge() {
	tests := []struct {
		intervals [][]int
		res       [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 4}, {2, 3}},
			[][]int{{1, 4}},
		},
		{
			[][]int{{1, 4}, {4, 5}, {5, 50}},
			[][]int{{1, 50}},
		},
		{
			[][]int{{1, 4}, {5, 8}},
			[][]int{{1, 4}, {5, 8}},
		},
	}

	for _, test := range tests {
		res := merge(test.intervals)
		fmt.Println(common.IntEqualEx(res, test.res))
	}
}
