package common

import "sort"

func StringEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func StringEqualEx(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	sort.Strings(str1)
	sort.Strings(str2)
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func StringArrayEqualEx(str1, str2 [][]string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if !StringEqualEx(str1[i], str2[i]) {
			return false
		}
	}
	return true
}
