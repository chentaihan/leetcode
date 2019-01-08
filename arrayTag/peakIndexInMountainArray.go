package arrayTag

import "fmt"

/**
852. 山脉数组的峰顶索引
我们把符合下列属性的数组 A 称作山脉：

A.length >= 3
存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。
分析：二分查找的变形
*/

func peakIndexInMountainArray(A []int) int {
	start, end := 0, len(A)-1
	for start <= end {
		middle := (start + end) / 2
		if A[middle] > A[middle+1] && A[middle] > A[middle-1] {
			return middle
		}
		if A[middle] < A[middle+1] {
			start = middle + 1
		} else if A[middle] > A[middle+1] {
			end = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

func TestpeakIndexInMountainArray() {
	tests := []struct {
		A   []int
		ret int
	}{
		{
			[]int{1, 3, 1},
			1,
		},
		{
			[]int{1, 3, 4, 1},
			2,
		},
		{
			[]int{1, 3, 4, 5, 1},
			3,
		},
		{
			[]int{1, 3, 4, 5, 4, 1},
			3,
		},
		{
			[]int{1, 3, 4, 5, 4, 3, 1},
			3,
		},
		{
			[]int{1, 3, 4, 5, 4, 3, 2, 1},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			4,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 4, 3, 2, 1},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1},
			5,
		},
	}

	for _, test := range tests {
		ret := peakIndexInMountainArray(test.A)
		fmt.Println(ret == test.ret)
	}
}
