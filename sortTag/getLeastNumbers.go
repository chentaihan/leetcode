package sortTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
面试题40. 最小的k个数
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

 

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
 

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	getLeastNumber(arr, k)
	return arr[:k]
}

func getLeastNumber(arr []int, k int) {
	index := getLeastNumbersPartion(arr)
	if index+1 == k {
		return
	}
	if index >= k {
		getLeastNumbers(arr[:index], k)
	} else {
		getLeastNumbers(arr[index+1:], k-index-1)
	}
}

func getLeastNumbersPartion(arr []int) int {
	flag, index := arr[0], 1
	start, end := 0, len(arr)-1
	for start < end {
		if arr[index] > flag {
			arr[index], arr[end] = arr[end], arr[index]
			end--
		} else {
			arr[index], arr[start] = arr[start], arr[index]
			start++
			index++
		}
	}
	arr[start] = flag
	return start
}

func TestgetLeastNumbers() {
	tests := []struct {
		arr []int
		k   int
		res []int
	}{
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			6,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			1,
			[]int{1},
		},
	}

	for _, test := range tests {
		res := getLeastNumbers(test.arr, test.k)
		fmt.Println(common.IntEqualSort(res, test.res))
	}
}
