package mapTag

import (
	"strconv"
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, len(cpdomains))
	for _, cpdomain := range cpdomains {
		count, strs := splitString(cpdomain)
		for _, str := range strs {
			m[str] += count
		}
	}
	ret := make([]string, len(m))
	index := 0
	for key, value := range m {
		ret[index] = strconv.Itoa(value) + key
		index++
	}
	return ret
}

func splitString(cpdoma string) (count int, strs []string) {
	i := 1
	for ; i < len(cpdoma); i++ {
		if cpdoma[i] == ' ' {
			count, _ = strconv.Atoi(cpdoma[:i])
			break
		}
	}

	strs = append(strs, cpdoma[i:])
	for i++; i < len(cpdoma); i++ {
		if cpdoma[i] == '.' {
			strs = append(strs, " "+cpdoma[i+1:])
			if len(strs) == 3 {
				break
			}
		}
	}
	return
}

func TestsubdomainVisits() {
	tests := []struct {
		cpdomains []string
		result    []string
	}{
		{
			[]string{"100 www.baidu.com"},
			[]string{"100 www.baidu.com", "100 baidu.com", "100 com"},
		},
		{
			[]string{"100 www.sina.com", "100 www.baidu.com"},
			[]string{"100 www.baidu.com", "100 www.sina.com", "100 baidu.com", "100 sina.com", "200 com"},
		},
		{
			[]string{"100 www.sina.com", "100 www.baidu.com", "100 abc.baidu.com"},
			[]string{"100 abc.baidu.com", "100 www.baidu.com", "100 www.sina.com", "200 baidu.com", "100 sina.com", "300 com"},
		},
	}
	for _, test := range tests {
		fmt.Println(common.StringEqualEx(subdomainVisits(test.cpdomains), test.result))
	}
}
