package mapTag

import "fmt"

/**
575. 分糖果
给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

示例 1:

输入: candies = [1,1,2,2,3,3]
输出: 3
解析: 一共有三种种类的糖果，每一种都有两个。
     最优分配方案：妹妹获得[1,2,3],弟弟也获得[1,2,3]。这样使妹妹获得糖果的种类数最多。
示例 2 :

输入: candies = [1,1,2,3]
输出: 2
解析: 妹妹获得糖果[2,3],弟弟获得糖果[1,1]，妹妹有两种不同的糖果，弟弟只有一种。这样使得妹妹可以获得的糖果种类数最多。
注意:

数组的长度为[2, 10,000]，并且确定为偶数。
数组中数字的大小在范围[-100,000, 100,000]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distribute-candies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func distributeCandies(candies []int) int {
	l := len(candies)
	m := map[int]bool{}
	for i := 0; i < l; i++ {
		m[candies[i]] = true
		if len(m) >= l/2 {
			return l/2
		}
	}
	return len(m)
}

func TestdistributeCandies() {
	tests := []struct {
		candies []int
		count   int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 1, 2, 2, 3, 3},
			3,
		},
		{
			[]int{1, 1, 1, 2, 3, 3},
			3,
		},
		{
			[]int{1, 1, 1, 1, 2, 3},
			3,
		},
		{
			[]int{1, 1, 1, 1, 1, 3},
			2,
		},
		{
			[]int{1, 1, 1, 1, 1, 1},
			1,
		},
	}
	for _, test := range tests {
		count := distributeCandies(test.candies)
		fmt.Println(count == test.count)
	}
}
