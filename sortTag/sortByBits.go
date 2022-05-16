package sortTag

import "sort"

/**
1356. 根据数字二进制下 1 的数目排序
给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。

如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。

请你返回排序后的数组。



示例 1：

输入：arr = [0,1,2,3,4,5,6,7,8]
输出：[0,1,2,4,8,3,5,6,7]
解释：[0] 是唯一一个有 0 个 1 的数。
[1,2,4,8] 都有 1 个 1 。
[3,5,6] 有 2 个 1 。
[7] 有 3 个 1 。
按照 1 的个数排序得到的结果数组为 [0,1,2,4,8,3,5,6,7]
示例 2：

输入：arr = [1024,512,256,128,64,32,16,8,4,2,1]
输出：[1,2,4,8,16,32,64,128,256,512,1024]
解释：数组中所有整数二进制下都只有 1 个 1 ，所以你需要按照数值大小将它们排序。
示例 3：

输入：arr = [10000,10000]
输出：[10000,10000]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sortByBits(arr []int) []int {
	list := make([]item, len(arr))
	for i := 0; i < len(list); i++ {
		list[i] = item{
			arr[i], count1(arr[i]),
		}
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].index == list[j].index {
			return list[i].val < list[j].val
		}
		return list[i].index < list[j].index
	})
	result := make([]int, len(arr))
	for i := 0; i < len(list); i++ {
		result[i] = list[i].val
	}
	return result
}

type item struct {
	val   int
	index int
}

func count1(x int) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
