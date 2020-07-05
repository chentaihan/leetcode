package stringTag

import (
	"fmt"
	"strconv"
	"strings"
)

//394. 字符串解码
//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
// 示例:
//
//
//s = "3[a]2[bc]", 返回 "aaabcbc".
//s = "3[a2[c]]", 返回 "accaccacc".
//s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
//
//

func decodeString(s string) string {
	if strings.Index(s, "[") < 0 {
		return s
	}
	var result []byte
	for len(s) > 0 {
		var buf []byte
		buf, s = decodeParse(s)
		val := decodeValue(string(buf))
		result = append(result, val...)
	}
	return string(result)
}

func decodeValue(buf string) []byte {
	if len(buf) == 0 {
		return []byte{}
	}
	if decodeJudge(buf) {
		index := 1
		for ; index < len(buf); index++ {
			if buf[index] < '0' || buf[index] > '9' {
				break
			}
		}
		count, _ := strconv.Atoi(buf[:index])
		value := decodeValue(buf[index+1 : len(buf)-1])
		result := make([]byte, count*len(value))
		for i := 0; i < count*len(value); i += len(value) {
			copy(result[i:], value)
		}
		return result
	} else {
		return []byte(decodeString(buf))
	}
}

func decodeJudge(buf string) bool {
	if buf[0] >= '0' && buf[0] <= '9' && buf[len(buf)-1] == ']' {
		count, i := 0, 0
		for ; i < len(buf); i++ {
			if buf[i] == '[' {
				count++
			} else if buf[i] == ']' {
				count--
				if count == 0 {
					break
				}
			}
		}
		if i == len(buf)-1 {
			return true
		}
	}
	return false
}

func decodeParse(s string) ([]byte, string) {
	if len(s) == 0 {
		return []byte{}, ""
	}
	var buf []byte
	if s[0] >= '0' && s[0] <= '9' {
		count, i := 0, 0
		for ; i < len(s); i++ {
			buf = append(buf, s[i])
			if s[i] == '[' {
				count++
			} else if s[i] == ']' {
				count--
				if count == 0 {
					break
				}
			}
		}
		return buf, s[i+1:]
	} else {
		i := 0
		for ; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				break
			}
			buf = append(buf, s[i])
		}
		return buf, s[i:]
	}
}

func TestdecodeString() {
	tests := []struct {
		s      string
		result string
	}{
		{
			"3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
			"zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
		},
		{
			"3[abc]",
			"abcabcabc",
		},
		{
			"abc",
			"abc",
		},

		{
			"3[abc]ef",
			"abcabcabcef",
		},
		{
			"3[a2[c]b]",
			"accbaccbaccb",
		},
		{
			"3[a2[c2[xx]]b]",
			"acxxxxcxxxxbacxxxxcxxxxbacxxxxcxxxxb",
		},
		{
			"3[a]2[bc]",
			"aaabcbc",
		},
		{
			"3[a2[c]]",
			"accaccacc",
		},
		{
			"2[abc]3[cd]ef",
			"abcabccdcdcdef",
		},
		{
			"2[abc]3[cd]ef2[abc]3[cd]",
			"abcabccdcdcdefabcabccdcdcd",
		},
	}

	for _, test := range tests {
		result := decodeString(test.s)
		if result != test.result {
			fmt.Println(result, "  ", test.result)
		} else {
			fmt.Println(true)
		}

	}
}
