package arrayTag

import (
	"fmt"
	"sort"
)

/**
1046. 最后一块石头的重量
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。

 

示例：

输入：[2,7,4,1,8,1]
输出：1
解释：
先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-stone-weight
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	l := len(stones)
	for l > 1 {
		left := stones[l-1] - stones[l-2]
		if left == 0 {
			l -= 2
			stones = stones[:l]
		} else {
			stones = stones[:l-2]
			index := binarySearchPos(stones, left)
			stones = IntInsertValue(stones, index, left)
			l -= 1
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

//二分查找插入的位置
func binarySearchPos(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1

		}
	}

	if target < nums[middle] {
		return middle
	} else {
		return middle + 1
	}
}

//在指定位置插入一个元素
func IntInsertValue(nums []int, index, value int) []int {
	nums = append(nums, value)
	if index >= len(nums) {
		return nums
	}
	if index < 0 {
		index = 0
	}
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = value
	return nums
}

func TestlastStoneWeight() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 3},
			2,
		},
		{
			[]int{1, 2, 3},
			0,
		},
		{
			[]int{2, 7, 4, 1, 8, 1},
			1,
		},
	}
	for index, test := range tests {
		result := lastStoneWeight(test.nums)
		fmt.Println(index, result == test.result)
	}
}
