package arrayTag

import "fmt"

/**
922. 按奇偶排序数组 II
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
你可以返回任何满足上述条件的数组作为答案。
 */

func sortArrayByParityII(A []int) []int {
	start, end := 0, len(A)-1
	for start < end {
		for (start%2)^(A[start]%2) == 0 && start < end {
			start++
		}
		for (end%2)^(A[end]%2) == 0 && end > 0 {
			end--
		}
		tmpEnd := end
		if tmpEnd > 0 {
			for (start %2) ^ (A[tmpEnd]%2) != 0 && tmpEnd > 0 {
				tmpEnd--
			}
			A[start], A[tmpEnd] = A[tmpEnd], A[start]
			start++
		}
	}
	return A
}

func checkSortArrayByParityII(A []int) bool {
	for index, val := range A {
		if (index%2)^(val%2) != 0 {
			return false
		}
	}
	return true
}

func TestsortArrayByParityII() {
	tests := []struct {
		A []int
	}{
		{
			[]int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3},
		},
		{
			[]int{2, 3, 1, 1, 4, 2, 2, 4, 3, 3},
		},
		{
			[]int{1, 2},
		},
		{
			[]int{2, 1},
		},
		{
			[]int{2, 1, 2, 1},
		},
		{
			[]int{1, 2, 1, 2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{2, 1, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{2, 1, 3, 4, 6, 5, 7, 8},
		},
		{
			[]int{2, 2, 3, 4, 6, 5, 7, 7},
		},
		{
			[]int{2, 2, 4, 3, 6, 5, 7, 7},
		},
		{
			[]int{4, 2, 5, 7},
		},
		{
			[]int{4, 5, 2, 7},
		},
		{
			[]int{5, 7, 2, 4},
		},
		{
			[]int{5, 2, 7, 4},
		},
	}

	for _, test := range tests {
		r := sortArrayByParityII(test.A)
		fmt.Println(checkSortArrayByParityII(r))
	}
}
