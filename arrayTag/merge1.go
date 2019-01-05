package arrayTag

import (
	"fmt"
	"sort"
)

/**
56. 合并区间 TODO
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

func (p IntervalSlice) Len() int { return len(p) }
func (p IntervalSlice) Less(i, j int) bool {
	if p[i].Start < p[j].Start {
		return true
	}
	return p[i].Start == p[j].Start && p[i].End > p[j].End
}
func (p IntervalSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func merge1(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(IntervalSlice(intervals))

	ret := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		cross := false
		for j := 0; j < len(ret); j++ {
			if isCross(&ret[j], &intervals[i]) {
				cross = true
				break
			}
		}
		if !cross {
			ret = append(ret, intervals[i])
		}
	}

	return ret
}

func isCross(i1, i2 *Interval) bool {
	if (i2.Start >= i1.Start && i2.Start <= i1.End) || (i2.End >= i1.Start && i2.End <= i1.End) {
		i1.Start = min(i1.Start, i2.Start)
		i1.End = max(i1.End, i2.End)
		return true
	}
	return false
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func Testmerge1() {
	tests := []struct {
		intervals []Interval
		result    []Interval
	}{
		{
			[]Interval{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}},
			[]Interval{{1, 3}, {4, 7}},
		},
		{
			[]Interval{{1, 3}, {2, 4}},
			[]Interval{{1, 4}},
		},
		{
			[]Interval{{1, 3}, {3, 6}},
			[]Interval{{1, 6}},
		},
		{
			[]Interval{{1, 3}, {4, 6}},
			[]Interval{{1, 3}, {4, 6}},
		},
		{
			[]Interval{{1, 3}, {4, 6}, {3, 5}},
			[]Interval{{1, 6}},
		},
		{
			[]Interval{{1, 3}, {3, 4}, {4, 6}, {5, 6}},
			[]Interval{{1, 6}},
		},
		{
			[]Interval{{1, 3}, {3, 4}, {4, 6}, {6, 7}},
			[]Interval{{1, 7}},
		},
	}

	for _, test := range tests {
		result := merge1(test.intervals)
		fmt.Println(IntervalsEqual(result, test.result))
	}
}

func IntervalsEqual(arr1, arr2 []Interval) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i].Start != arr2[i].Start || arr1[i].End != arr2[i].End {
			return false
		}
	}
	return true
}
