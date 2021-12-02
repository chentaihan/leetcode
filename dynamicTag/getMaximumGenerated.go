package dynamicTag

import "fmt"

/**
1646. 获取生成数组中的最大值
*/

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}
	nums := make([]int, n+1)
	nums[1] = 1
	max := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[(i+1)/2]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func TestgetMaximumGenerated() {
	tests := []struct {
		n   int
		res int
	}{
		{
			7, 3,
		},
		{
			3, 2,
		},
		{
			2, 1,
		},
	}
	for _, test := range tests {
		res := getMaximumGenerated(test.n)
		fmt.Println(res == test.res)
	}
}
