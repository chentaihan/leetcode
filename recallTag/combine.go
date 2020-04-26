package recallTag

import "fmt"

/**
77. 组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combine(n int, k int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	var result [][]int
	var list []int
	_combine(n, k, &result, &list)
	return result
}

func _combine(num, k int, result *[][]int, list *[]int) {
	if len(*list) == k {
		array := make([]int, k)
		copy(array, *list)
		*result = append(*result, array)
		return
	}
	for i := 1; i <= num; i++ {
		if !inArray(*list, i) {
			if len(*list) == 0 || (len(*list) > 0 && (*list)[len(*list)-1] < i) {
				*list = append(*list, i)
				_combine(num, k, result, list)
				*list = (*list)[:len(*list)-1]
			}
		}
	}
}

func combineEx(n int, k int) [][]int {
	l := make([]int, k)
	i := 0
	var ret [][]int
	for i >= 0 {
		l[i]++
		if l[i] > n-k+i+1 {
			i--
		} else if i == k-1 {
			t := make([]int, k)
			copy(t, l)
			ret = append(ret, t)
		} else {
			i++
			l[i] = l[i-1]
		}
	}
	return ret
}

func Testcombine() {
	tests := []struct {
		n int
		k int
	}{
		{
			1, 1,
		},
		{
			2, 2,
		},
		{
			3, 2,
		},
		{
			4, 2,
		},
	}

	for _, test := range tests {
		result := combine(test.n, test.k)
		fmt.Println(result)
	}
}
