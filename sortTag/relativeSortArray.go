package sortTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"sort"
)

/**
1122. 数组的相对排序
给你两个数组，arr1 和 arr2，

arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

 

示例：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]
 

提示：

arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i] 各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/relative-sort-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int, len(arr2))
	for i := 0; i < len(arr2); i++ {
		m[arr2[i]] = 1
	}

	var left []int
	for i := 0; i < len(arr1); i++ {
		if m[arr1[i]] >= 1 {
			m[arr1[i]]++
		} else {
			left = append(left, arr1[i])
		}
	}

	index := 0
	for i := 0; i < len(arr2); i++ {
		for j := m[arr2[i]]; j >= 2; j-- {
			arr1[index] = arr2[i]
			index++
		}
	}
	sort.Ints(left)
	copy(arr1[index:], left)
	return arr1
}

func TestrelativeSortArray() {
	tests := []struct {
		arr1 []int
		arr2 []int
		res  []int
	}{
		{
			[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			[]int{2, 1, 4, 3, 9, 6},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 1, 2, 3},
			[]int{1, 2, 3},
			[]int{1, 1, 2, 2, 3, 3, 4, 5, 6},
		},
	}
	for _, test := range tests {
		res := relativeSortArray(test.arr1, test.arr2)
		fmt.Println(common.IntEqual(res, test.res))
	}
}
