package arrayTag

import "sort"

/**
1051. 高度检查器
学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。

请你返回至少有多少个学生没有站在正确位置数量。该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。

示例：

输入：[1,1,4,2,1,3]
输出：3
解释：
高度为 4、3 和最后一个 1 的学生，没有站在正确的位置。


提示：

1 <= heights.length <= 100
1 <= heights[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/height-checker
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func heightChecker(heights []int) int {
	h := make([]int, len(heights))
	copy(h, heights)
	sort.Ints(h)
	count := 0
	for i := 0; i < len(h); i++ {
		if h[i] != heights[i] {
			count++
		}
	}
	return count
}
