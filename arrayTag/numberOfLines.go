package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
806. 写字符串需要的行数
我们要把给定的字符串 S 从左到右写到每一行上，每一行的最大宽度为100个单位，如果我们在写某个字母的时候会使这行超过了100 个单位，那么我们应该把这个字母写到下一行。我们给定了一个数组 widths ，这个数组 widths[0] 代表 'a' 需要的单位， widths[1] 代表 'b' 需要的单位，...， widths[25] 代表 'z' 需要的单位。

现在回答两个问题：至少多少行能放下S，以及最后一行使用的宽度是多少个单位？将你的答案作为长度为2的整数列表返回。

示例 1:
输入:
widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "abcdefghijklmnopqrstuvwxyz"
输出: [3, 60]
解释:
所有的字符拥有相同的占用单位10。所以书写所有的26个字母，
我们需要2个整行和占用60个单位的一行。
示例 2:
输入:
widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "bbbcccdddaaa"
输出: [2, 4]
解释:
除去字母'a'所有的字符都是相同的单位10，并且字符串 "bbbcccdddaa" 将会覆盖 9 * 10 + 2 * 4 = 98 个单位.
最后一个字母 'a' 将会被写到第二行，因为第一行只剩下2个单位了。
所以，这个答案是2行，第二行有4个单位宽度。


注:

字符串 S 的长度在 [1, 1000] 的范围。
S 只包含小写字母。
widths 是长度为 26的数组。
widths[i] 值的范围在 [2, 10]。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/number-of-lines-to-write-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numberOfLines(widths []int, s string) []int {
	lines, total := 0, 0
	for i := 0; i < len(s); i++ {
		val := widths[s[i]-'a']
		sum := total + val
		if sum > 100 {
			lines++
			total = val
		} else if sum == 100 {
			lines++
			total = 0
		} else {
			total += val
		}
	}
	if total > 0 {
		lines++
	} else {
		total = 100
	}
	return []int{lines, total}
}

func TestNumberOfLines() {
	tests := []struct {
		widths []int
		s      string
		res    []int
	}{
		{
			[]int{3, 4, 10, 4, 8, 7, 3, 3, 4, 9, 8, 2, 9, 6, 2, 8, 4, 9, 9, 10, 2, 4, 9, 10, 8, 2},
			"mqblbtpvicqhbrejb",
			[]int{1, 100},
		},
		{
			[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"abcdefghijklmnopqrstuvwxyz",
			[]int{3, 60},
		},
		{
			[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"bbbcccdddaaa",
			[]int{2, 4},
		},
	}
	for _, test := range tests {
		res := numberOfLines(test.widths, test.s)
		if common.IntEqual(res, test.res) {
			fmt.Println(true)
		} else {
			fmt.Println(res, test.res)
		}
	}
}
