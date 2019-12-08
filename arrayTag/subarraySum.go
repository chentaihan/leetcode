package arrayTag

import "fmt"

/**
560. 和为K的子数组
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
TODO
*/

func subarraySum(nums []int, k int) int {
	sum, start, count := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < k {
			if nums[i] > k {
				for sum < k && start < i {
					sum -= nums[start]
					start++
					if sum == k {
						count++
					}
				}
			} else if nums[i] == k {
				count++
				start = i
				sum = k
			}
			continue
		}
		if sum == k {
			count++
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 0 {
					count++
				} else {
					break
				}
			}
		}
		for sum >= k && start < i {
			sum -= nums[start]
			start++
			if sum == k {
				count++
			}
		}
	}
	return count
}

func TestsubarraySum() {
	tests := []struct {
		nums  []int
		k     int
		count int
	}{
		{
			[]int{28,54,7,-70,22,65,-6},
			100,
			1,
		},
		{
			[]int{-1, -1, 1},
			0,
			1,
		},
		{
			[]int{-1, -1, 1, 1},
			1,
			2,
		},
		{
			[]int{-1, -1, 1, 1, 0},
			1,
			3,
		},
		{
			[]int{-1, -1, 1},
			1,
			1,
		},
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{},
			1,
			0,
		},
		{
			[]int{1},
			1,
			1,
		},
		{
			[]int{1, 1, 1},
			1,
			3,
		},
		{
			[]int{1, 1, 1},
			2,
			2,
		},
		{
			[]int{1, 2, 1, 2},
			3,
			3,
		},
		{
			[]int{1, 2, 0, 1, 2},
			3,
			5,
		},
		{
			[]int{1, 2, 0, 0, 1, 2},
			3,
			7,
		},
		{
			[]int{1, 2, 1, 2, 1},
			3,
			4,
		},
		{
			[]int{1, 2, 1, 2, 1, 2},
			3,
			5,
		},
		{
			[]int{1, 2, 4, 6, 8, 4, 2, 1},
			12,
			2,
		},
		{
			[]int{1, 2, 4, 6, 8, 4, 2, 1, 5},
			12,
			3,
		},
		{
			[]int{1, 2, 4, 6, 8, 4, 2, 1, 5, 4},
			12,
			4,
		},
		{
			[]int{1, 2, 4, 6, 8, 4, 2, 1, 5, 4, 2},
			12,
			5,
		},
	}
	for _, test := range tests {
		count := subarraySum(test.nums, test.k)
		fmt.Println(count == test.count)
	}
}
