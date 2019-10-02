package stringTag

import (
	"bytes"
	"fmt"
	"strconv"
)

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

func decodeStringold(s string) string {
	var buffer bytes.Buffer
	var countList []int
	in := false
	var inBuf []byte
	for i := 0; i < len(s); i++ {
		if in {
			if s[i] == ']' {
				in = false
				count := countList[len(countList)-1]
				countList = countList[:len(countList)-1]
				for count > 0 {
					buffer.Write(inBuf)
					count--
				}
				inBuf = inBuf[:0]
			} else {
				inBuf = append(inBuf, s[i])
			}
			continue
		}
		if s[i] == '[' {
			in = true
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			start := i
			for s[i] >= '0' && s[i] <= '9' {
				i++
			}
			count, _ := strconv.Atoi(string(s[start:i]))
			countList = append(countList, count)
			i--
		} else {
			buffer.WriteByte(s[i])
		}
	}
	return buffer.String()
}

type token struct {
	before  string
	after   string
	count   int
	content string
}

type tokenNode struct {
	token
	children []*tokenNode
}

func (tn *tokenNode) String() string {
	var buffer bytes.Buffer
	if len(tn.children) == 0 {
		buffer.WriteString(tn.before)
		for i := 0; i < tn.count; i++ {
			buffer.WriteString(tn.content)
		}
		buffer.WriteString(tn.after)
		return buffer.String()
	}
	buffer.WriteString(tn.before)
	for i := 0; i < tn.count; i++ {
		for _, child := range tn.children {
			buffer.WriteString(child.String())
		}
	}
	buffer.WriteString(tn.after)
	return buffer.String()
}

func decodeString(s string) string {
	nodeList := []*tokenNode{
		&tokenNode{},
	}
	index := 0
	curNode := nodeList[index]

	var buffer bytes.Buffer
	depth := 0
	start := 0
	before, after := "", ""
	var inbuf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			depth--
			start = i + 1
			inbuf.Reset()
			continue
		}
		if s[i] == '[' {
			depth++
			start = i + 1
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			if before == "" {
				before = s[start:i]
			} else {
				after = s[start:i]
			}

			start := i
			for s[i] >= '0' && s[i] <= '9' {
				i++
			}
			count, _ := strconv.Atoi(string(s[start:i]))
			i--
			if depth < 1 {
				curNode.before = before
				curNode.count = count
				curNode.after = after
			} else {
				node := &tokenNode{
					token: token{
						before: before,
						count:  count,
						after:  after,
					},
				}
				curNode.children = append(curNode.children, node)
				curNode = node
			}
			before, after = "", ""
		} else {
			inbuf.WriteByte(s[i])
		}
	}
	return buffer.String()
}

func TestdecodeString() {
	tests := []struct {
		s      string
		result string
	}{
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
			"dssdsd3[a2[c]bb3[cc]gff]dsdsd",
			"accaccacc",
		},
	}

	for _, test := range tests {
		result := decodeString(test.s)
		fmt.Println(result == test.result)
	}
}
