package mapTag

import "fmt"

/**
961. 重复 N 次的元素
在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。
返回重复了 N 次的那个元素。
 */

func repeatedNTimes(A []int) int {
	m := make(map[int]int)
	count := len(A) / 2
	for i := 0; i < len(A); i++ {
		m[A[i]]++
		if m[A[i]] == count {
			return A[i]
		}
	}
	return 0
}

func TestrepeatedNTimes() {
	tests := []struct{
		A []int
		ret int
	} {
		{
			[]int{1,2,3,4,2,2},
			2,
		},
		{
			[]int{1,2,3,2},
			2,
		},
		{
			[]int{1,2,3,2,3,3},
			3,
		},
		{
			[]int{1,2,2,3,3,3},
			3,
		},
	}

	for _,test := range tests {
		ret := repeatedNTimes(test.A)
		fmt.Println(ret == test.ret)
	}
}