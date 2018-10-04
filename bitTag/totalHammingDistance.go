package bitTag

import (
	"fmt"
)

/**
477. 汉明距离总和
统计二进制每个位的个数：total = count*(len(nums)-count)
 */

func totalHammingDistance(nums []int) int {
	total := 0
	l := len(nums)
	for index := 0; index < 63; index++ {
		count := 0
		for i := 0; i < l; i++ {
			if nums[i]&1 == 1 {
				count++
			}
			nums[i] >>= 1
		}
		total += count * (l - count)
	}
	return total
}

//超时
func totalHammingDistanceV1(nums []int) int {
	total := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			total += hammingDistance(nums[i], nums[j])
		}
	}
	return total
}

//超时
func totalHammingDistanceV2(nums []int) int {
	total := 0
	l := len(nums)
	buffer := make([][]int, l)
	for i := 0; i < l; i++ {
		buffer[i] = calBitCount(nums[i])
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			total += compareBit(buffer[i], buffer[j])
		}
	}
	return total
}

func calBitCount(x int) []int {
	result := make([]int, 0, 8)
	for i := 0; x > 0; i++ {
		if x&1 == 1 {
			result = append(result, i)
		}
		x >>= 1
	}
	return result
}

func compareBit(x, y []int) int {
	xIndex, yIndex := 0, 0
	count := 0
	for xIndex < len(x) && yIndex < len(y) {
		if x[xIndex] == y[yIndex] {
			xIndex++
			yIndex++
		} else if x[xIndex] < y[yIndex] {
			xIndex++
			count++
		} else {
			yIndex++
			count++
		}
	}
	count += len(x) - xIndex + len(y) - yIndex
	return count
}

func TesttotalHammingDistance() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{4, 14, 2},
			6,
		},
		{
			[]int{7, 8, 15},
			8,
		},
		{
			[]int{7, 8, 16},
			10,
		},
		{
			[]int{4, 8, 16},
			6,
		},
		{
			[]int{7, 56, 448},
			18,
		},
	}

	for _, test := range tests {
		result := totalHammingDistance(test.nums)
		fmt.Println(result == test.result)
	}
}
