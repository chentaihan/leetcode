package title

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}
	if digits[0] == 0 {
		newDigits := make([]int, len(digits)+1)
		copy(newDigits[1:], digits)
		newDigits[0] = 1
		return newDigits
	}
	return digits
}

func TestPlusOne() {
	tests := []struct {
		digits1 []int
		digits2 []int
	}{
		{
			nil,
			nil,
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{0},
			[]int{1},
		},
		{
			[]int{5},
			[]int{6},
		},
		{
			[]int{9},
			[]int{1, 0},
		},
		{
			[]int{1, 0},
			[]int{1, 1},
		},
		{
			[]int{1, 8},
			[]int{1, 9},
		},
		{
			[]int{1, 9},
			[]int{2, 0},
		},
		{
			[]int{9, 9},
			[]int{1, 0, 0},
		},
		{
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
		},
		{
			[]int{9, 9, 9, 9},
			[]int{1, 0, 0, 0, 0},
		},
		{
			[]int{9, 9, 9, 9, 9},
			[]int{1, 0, 0, 0, 0, 0},
		},
		{
			[]int{2, 9, 9, 9, 9, 9},
			[]int{3, 0, 0, 0, 0, 0},
		},
		{
			[]int{2, 9, 9, 9, 9, 9, 8},
			[]int{2, 9, 9, 9, 9, 9, 9},
		},
		{
			[]int{2, 9, 9, 9, 9, 9, 1},
			[]int{2, 9, 9, 9, 9, 9, 2},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1},
		},
	}
	for _, test := range tests {
		fmt.Println(intArrayEqual(plusOne(test.digits1), test.digits2))
	}
}
