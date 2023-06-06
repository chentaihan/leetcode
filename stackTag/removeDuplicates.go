package stackTag

import "fmt"

/**
1047. 删除字符串中的所有相邻重复项
给出由小写字母组成的字符串S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。



示例：

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。


提示：

1 <= S.length <= 20000
S 仅由小写英文字母组成。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func TestRemoveDuplicates() {
	tests := []struct {
		s   string
		res string
	}{
		{
			"abbaca",
			"ca",
		},
		{
			"aabbcc",
			"",
		},
		{
			"aabbccc",
			"c",
		},
		{
			"aaabbccc",
			"ac",
		},
		{
			"aaabbcc",
			"a",
		},
	}
	for _, test := range tests {
		res := removeDuplicates(test.s)
		if res == test.res {
			fmt.Println(true)
		} else {
			fmt.Println(test.res, res)
		}
	}
}
