package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
1053. 交换一次的先前排列
给你一个正整数的数组 A（其中的元素不一定完全不同），请你返回可在 一次交换（交换两数字 A[i] 和 A[j] 的位置）后得到的、按字典序排列小于 A 的最大可能排列。

如果无法这么操作，就请返回原数组。

 

示例 1：

输入：arr = [3,2,1]
输出：[3,1,2]
解释：交换 2 和 1
示例 2：

输入：arr = [1,1,5]
输出：[1,1,5]
解释：已经是最小排列
示例 3：

输入：arr = [1,9,4,6,7]
输出：[1,7,4,6,9]
解释：交换 9 和 7
示例 4：

输入：arr = [3,1,1,3]
输出：[1,3,1,3]
解释：交换 1 和 3
 

提示：

1 <= arr.length <= 104
1 <= arr[i] <= 104


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/previous-permutation-with-one-swap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：从右往左扫描，找到第一个左值大于右值的左值为left，从这个left往右扫，找到<arr[left]的最大值
*/

func prevPermOpt1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := -1
	for i := len(arr) - 1; i >= 1; i-- {
		if arr[i-1] > arr[i] {
			left = i - 1
			break
		}
	}
	if left == -1 {
		return arr
	}
	right := left + 1
	for i := left + 2; i < len(arr); i++ {
		if arr[i] > arr[right] && arr[i] < arr[left] {
			right = i
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	return arr
}

func TestprevPermOpt1() {
	tests := []struct {
		arr []int
		res []int
	}{
		{
			[]int{3, 2, 1},
			[]int{3, 1, 2},
		},
		{
			[]int{1, 1, 5},
			[]int{1, 1, 5},
		},
		{
			[]int{1, 9, 4, 6, 7},
			[]int{1, 7, 4, 6, 9},
		},
		{
			[]int{3, 1, 1, 3},
			[]int{1, 3, 1, 3},
		},
		{
			[]int{3, 1, 2, 3},
			[]int{2, 1, 3, 3},
		},
		{
			[]int{4, 1, 1, 3},
			[]int{3, 1, 1, 4},
		},
	}
	for index, test := range tests {
		res := prevPermOpt1(test.arr)
		if !common.IntEqual(res, test.res) {
			fmt.Println(false, index, test.res, res)
		} else {
			fmt.Println(true)
		}

	}
}
