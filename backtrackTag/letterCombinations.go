package backtrackTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。


示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	result := make([]string, 0, len(digits))
	list := make([]byte, 0, len(digits))
	_letterCombinations(m, &result, list, digits, 0)
	return result
}

func _letterCombinations(m []string, result *[]string, list []byte, digits string, index int) {
	if index == len(digits) {
		*result = append(*result, string(list))
		return
	}
	curLine := m[digits[index]-'2']
	for i := 0; i < len(curLine); i++ {
		list = append(list, curLine[i])
		_letterCombinations(m, result, list, digits, index+1)
		list = list[:len(list)-1]
	}
}

func TestletterCombinations() {
	tests := []struct {
		digits string
		result []string
	}{
		{
			"",
			[]string{},
		},
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"24",
			[]string{"ag", "ah", "ai", "bg", "bh", "bi", "cg", "ch", "ci"},
		},
		{ // pqrs
			"27",
			[]string{"ap", "aq", "ar", "as", "bp", "bq", "br", "bs", "cp", "cq", "cr", "cs"},
		},
		{
			"2",
			[]string{"a", "b", "c"},
		},
		{
			"3",
			[]string{"d", "e", "f"},
		},
	}
	for _, test := range tests {
		result := letterCombinations(test.digits)
		fmt.Println(common.StringEqualEx(result, test.result))
	}
}
