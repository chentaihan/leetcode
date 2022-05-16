package arrayTag

import "fmt"

/**
532. 数组中的K-diff数对
给定一个整数数组和一个整数 k, 你需要在数组里找到不同的 k-diff 数对。这里将 k-diff 数对定义为一个整数对 (i, j), 其中 i 和 j 都是数组中的数字，且两数之差的绝对值是 k.

示例 1:

输入: [3, 1, 4, 1, 5], k = 2
输出: 2
解释: 数组中有两个 2-diff 数对, (1, 3) 和 (3, 5)。
尽管数组中有两个1，但我们只应返回不同的数对的数量。
示例 2:

输入:[1, 2, 3, 4, 5], k = 1
输出: 4
解释: 数组中有四个 1-diff 数对, (1, 2), (2, 3), (3, 4) 和 (4, 5)。
示例 3:

输入: [1, 3, 1, 5, 4], k = 0
输出: 1
解释: 数组中只有一个 0-diff 数对，(1, 1)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/k-diff-pairs-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	m := map[int]int{}
	count := 0
	if k == 0 {
		for _, val := range nums {
			if m[val] == 1 {
				count++
			}
			m[val]++
		}
	} else {
		for _, val := range nums {
			if m[val] > 0 {
				continue
			}
			m[val] = 1
			if m[val-k] > 0 {
				count++
			}
			if m[k+val] > 0 {
				count++
			}
		}
	}
	return count
}

func TestfindPairs() {
	tests := []struct {
		nums  []int
		k     int
		count int
	}{
		{
			[]int{1, 3, 5, 7, 9},
			2,
			4,
		},
		{
			[]int{1, 3, 5, 7, 9, 1},
			2,
			4,
		},
		{
			[]int{1, 3, 5, 7, 9, 1},
			2,
			4,
		},
		{
			[]int{1, 3, 1, 7, 9, 1},
			0,
			1,
		},
		{
			[]int{3, 1, 4, 1, 5},
			2,
			2,
		},
	}
	for _, test := range tests {
		count := findPairs(test.nums, test.k)
		fmt.Println(count == test.count)
	}
}
