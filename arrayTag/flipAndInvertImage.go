package arrayTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
832. 翻转图像
给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。

水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。

反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。
 */

func flipAndInvertImage(A [][]int) [][]int {
	val := []int{1, 0}
	for i := 0; i < len(A); i++ {
		curRow := A[i]
		start, end := 0, len(curRow)-1
		for start <= end {
			if curRow[start] == curRow[end] {
				curRow[start], curRow[end] = val[curRow[start]], val[curRow[start]]
			}
			start++
			end--
		}
	}
	return A
}

func TestflipAndInvertImage() {
	tests := []struct {
		A [][]int
		B [][]int
	}{
		{
			[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, test := range tests {
		B := flipAndInvertImage(test.A)
		for i := 0; i < len(B); i++ {
			fmt.Println(common.IntEqual(B[i], test.B[i]))
		}
	}
}
