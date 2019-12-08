package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"strconv"
)

/**
228. 汇总区间
给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。

示例 1:

输入: [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
示例 2:

输入: [0,2,3,4,6,8,9]
输出: ["0","2->4","6","8->9"]
解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/summary-ranges
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func summaryRanges(nums []int) []string {
	var result []string
	if len(nums) == 0 {
		return result
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		if index == i-1 {
			result = append(result, strconv.Itoa(nums[i-1]))
		} else {
			result = append(result, fmt.Sprintf("%v->%v", nums[index], nums[i-1]))
		}
		index = i
	}

	if index == len(nums)-1 {
		result = append(result, strconv.Itoa(nums[len(nums)-1]))
	} else {
		result = append(result, fmt.Sprintf("%v->%v", nums[index], nums[len(nums)-1]))
	}

	return result
}

func TestsummaryRanges() {
	tests := []struct {
		nums   []int
		reuslt []string
	}{
		{
			[]int{1},
			[]string{"1"},
		},
		{
			[]int{1, 2},
			[]string{"1->2"},
		},
		{
			[]int{1, 2, 3},
			[]string{"1->3"},
		},
		{
			[]int{1, 2, 3, 5},
			[]string{"1->3", "5"},
		},
		{
			[]int{1, 2, 3, 5, 6},
			[]string{"1->3", "5->6"},
		},
		{
			[]int{1, 2, 3, 5, 6, 7},
			[]string{"1->3", "5->7"},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8},
			[]string{"1->3", "5->8"},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 10},
			[]string{"1->3", "5->8", "10"},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 10, 12},
			[]string{"1->3", "5->8", "10", "12"},
		},
		{
			[]int{1, 6, 10, 12, 14},
			[]string{"1", "6", "10", "12", "14"},
		},
	}
	for _, test := range tests {
		result := summaryRanges(test.nums)
		fmt.Println(common.StringEqual(result, test.reuslt))
	}
}
