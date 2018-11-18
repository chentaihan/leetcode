package arrayTag

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

func merge1(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	ret := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {

	}
	return ret
}
