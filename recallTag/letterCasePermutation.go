package recallTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
784. 字母大小写全排列
分析：可以理解成找出二叉树所有从根节点到叶子节点的路径，字母为两个子节点，数字为1个子节点
 */

func letterCasePermutation(S string) []string {
	result := make([]string, 0, len(S))
	_letterCasePermutation([]byte(S), 0, &result)
	return result
}

func _letterCasePermutation(s []byte, index int, result *[]string) {
	if index >= len(s) {
		*result = append(*result, string(s))
		return
	}
	_letterCasePermutation(s, index+1, result)
	if (s[index] >= 'a' && s[index] <= 'z') || (s[index] >= 'A' && s[index] <= 'Z') {
		s[index] ^= 32
		_letterCasePermutation(s, index+1, result)
	}
}

func TestletterCasePermutation(){
	tests := []struct{
		S string
		result []string
	}{
		{
			"a1b2",
			[]string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			"3z4",
			[]string{"3z4", "3Z4"},
		},
	}

	for _,test := range tests{
		result := letterCasePermutation(test.S)
		fmt.Println(common.StringEqualEx(result, test.result))
	}
}
