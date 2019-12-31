package mapTag

import (
	"fmt"
	"sort"
)

/**
846. 一手顺子
爱丽丝有一手（hand）由整数数组给定的牌。 

现在她想把牌重新排列成组，使得每个组的大小都是 W，且由 W 张连续的牌组成。

如果她可以完成分组就返回 true，否则返回 false。

示例 1：

输入：hand = [1,2,3,6,2,3,4,7,8], W = 3
输出：true
解释：爱丽丝的手牌可以被重新排列为 [1,2,3]，[2,3,4]，[6,7,8]。
示例 2：

输入：hand = [1,2,3,4,5], W = 4
输出：false
解释：爱丽丝的手牌无法被重新排列成几个大小为 4 的组。
 

提示：

1 <= hand.length <= 10000
0 <= hand[i] <= 10^9
1 <= W <= hand.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hand-of-straights
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}
	m := make(map[int]int)
	for _, h := range hand {
		m[h]++
	}
	index := 0
	list := make([]int, len(m))
	for key := range m {
		list[index] = key
		index++
	}
	sort.Ints(list)

	for i := 0; i < len(list); {
		count := m[list[i]]
		if count == 0 {
			return false
		}
		k := i + W
		m[list[i]] -= count
		for j := i + 1; j < i+W; j++ {
			if j >= len(list) {
				return false
			}
			m[list[j]] -= count
			if j >= len(list) || list[j] != list[j-1]+1 || m[list[j]] < m[list[j-1]] {
				return false
			}
			if k == i+W && m[list[j]] > m[list[j-1]] {
				k = j
			}
		}
		i = k
	}
	return true
}

func TestisNStraightHand() {
	tests := []struct {
		hand   []int
		W      int
		result bool
	}{
		{
			[]int{1, 2, 3, 1, 2, 3},
			2,
			false,
		},
		{
			[]int{1},
			1,
			true,
		},
		{
			[]int{1, 2, 3, 5, 6, 7},
			1,
			true,
		},
		{
			[]int{1, 2, 3, 2, 3, 4, 3, 4, 5, 4, 5, 6, 5, 6, 7},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 5, 6, 7},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 2, 3, 4},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 1, 2, 3, 4},
			3,
			false,
		},
		{
			[]int{1, 2, 3, 1, 2, 3, 4, 6, 8},
			3,
			false,
		},
	}
	for _, test := range tests {
		result := isNStraightHand(test.hand, test.W)
		fmt.Println(result == test.result)
	}
}
