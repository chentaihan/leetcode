package arrayTag

/**
307. 区域和检索 - 数组可修改
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。

示例:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
说明:

数组仅可以在 update 函数下进行修改。
你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-mutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type NumArray struct {
	list []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		list: nums,
	}
}

func (this *NumArray) Update(i int, val int) {
	if len(this.list) < i || i < 0 {
		return
	}
	this.list[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 {
		i = 0
	}
	if j >= len(this.list) {
		j = len(this.list) - 1
	}
	sum := 0
	for ; i <= j; i++ {
		sum += this.list[i]
	}
	return sum
}
