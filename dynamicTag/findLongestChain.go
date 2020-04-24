package dynamicTag

import (
	"fmt"
	"sort"
)

/**
646. 最长数对链
给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。

现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。

给定一个对数集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。

示例 :

输入: [[1,2], [2,3], [3,4]]
输出: 2
解释: 最长的数对链是 [1,2] -> [3,4]
*/

func findLongestChain(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1]) {
			return true
		}
		return false
	})
	count := 1
	index := 0
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > pairs[index][1] {
			index = i
			count++
		} else if pairs[i][1] < pairs[index][1] {
			index = i
		}
	}
	return count
}

func TestfindLongestChain() {
	tests := []struct {
		pairs  [][]int
		result int
	}{
		{
			[][]int{{-10, -8}, {-6, -4}, {-5, 0}, {-4, 7}, {1, 7}, {6, 10}, {8, 9}, {9, 10}},
			4,
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}},
			2,
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			3,
		},
		{
			[][]int{{1, 2}, {3, 5}, {5, 6}},
			2,
		},
		{
			[][]int{{1, 2}, {3, 4}, {3, 6}, {5, 6}},
			3,
		},
	}

	for _, test := range tests {
		result := findLongestChain(test.pairs)
		fmt.Println(result == test.result)
	}
}
