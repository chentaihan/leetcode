package sortTag

import (
	"bytes"
	"fmt"
	"sort"
)

/**
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

type bytesSlice [][]byte

func (p bytesSlice) Len() int { return len(p) }
func (p bytesSlice) Less(i, j int) bool {
	iIndex, jIndex := 0, 0
	for iIndex < len(p[i]) && jIndex < len(p[j]) {
		if p[i][iIndex] < p[j][jIndex] {
			return false
		} else if p[i][iIndex] > p[j][jIndex] {
			return true
		}
		iIndex++
		jIndex++
	}
	if len(p[i]) < len(p[j]) && p[j][0] < p[j][jIndex] {
		return false
	}
	if len(p[i]) > len(p[j]) && p[i][iIndex] >= p[i][0] {
		return false
	}
	return true
}
func (p bytesSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func largestNumber(nums []int) string {
	var buf bytes.Buffer
	list := intToString(nums)
	sort.Sort(bytesSlice(list))
	for _, item := range list {
		buf.Write(item)
	}
	return buf.String()
}

func intToString(nums []int) [][]byte {
	result := make([][]byte, len(nums))
	buf := make([]byte, 20)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = []byte{'0'}
			continue
		}
		index := 0
		for nums[i] > 0 {
			val := nums[i] % 10
			buf[index] = byte('0' + val)
			nums[i] = nums[i] / 10
			index++
		}
		tBuf := make([]byte, index)
		k := 0
		for j := index - 1; j >= 0; j-- {
			tBuf[k] = buf[j]
			k++
		}
		result[i] = tBuf
	}
	return result
}

func TestlargestNumber() {
	tests := []struct {
		nums []int
		ret  string
	}{
		{
			[]int{99, 9, 98, 97, 96, 88, 97},
			"9999897979688",
		},
		{
			[]int{998, 9, 98, 97, 96, 88, 97},
			"99989897979688",
		},
		{
			[]int{99, 9, 98, 97, 66, 65, 644},
			"99998976665644",
		},
		{
			[]int{3, 30, 34, 5, 9},
			"9534330",
		},
		{
			[]int{12, 121},
			"12121",
		},
	}
	for _, test := range tests {
		result := largestNumber(test.nums)
		fmt.Println(result == test.ret)
	}
}
