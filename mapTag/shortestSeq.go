package mapTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"math"
)

/**
面试题 17.18. 最短超串
假设你有两个数组，一个长一个短，短的元素均不相同。找到长数组中包含短数组所有的元素的最短子数组，其出现顺序无关紧要。

返回最短子数组的左端点和右端点，如有多个满足条件的子数组，返回左端点最小的一个。若不存在，返回空数组。

示例 1:

输入:
big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7]
small = [1,5,9]
输出: [7,10]
示例 2:

输入:
big = [1,2,3]
small = [4]
输出: []
提示：

big.length <= 100000
1 <= small.length <= 100000


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-supersequence-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func shortestSeq(big []int, small []int) []int {
	mCheck := make(map[int]bool, len(small))
	m := make(map[int]int, len(big))
	for i := 0; i < len(small); i++ {
		mCheck[small[i]] = true
	}
	var res []int
	for i := 0; i < len(big); i++ {
		if mCheck[big[i]] {
			m[big[i]] = i
			if len(m) == len(mCheck) {
				list := _shortestSeq(m)
				if len(res) == 0 || list[1]-list[0] < res[1]-res[0] {
					res = list
				} else if list[1]-list[0] == res[1]-res[0] && list[0] < res[0] {
					res = list
				}
			}
		}
	}
	return res
}

func _shortestSeq(m map[int]int) []int {
	min, max := math.MaxInt64, math.MinInt64
	for _, val := range m {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return []int{min, max}
}

func TestshortestSeq() {
	tests := []struct {
		big   []int
		small []int
		res   []int
	}{
		{
			[]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7},
			[]int{1, 5, 9},
			[]int{7, 10},
		},
		{
			[]int{1, 2, 3},
			[]int{4},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{2},
			[]int{1, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 3},
			[]int{0, 2},
		},
	}
	for _, test := range tests {
		res := shortestSeq(test.big, test.small)
		fmt.Println(common.IntEqual(res, test.res))
	}
}
