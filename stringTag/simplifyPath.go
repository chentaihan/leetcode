package stringTag

import "fmt"

/**
71. 简化路径
给定一个文档 (Unix-style) 的完全路径，请进行路径简化。

例如，
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

边界情况:

你是否考虑了 路径 = "/../" 的情况？
在这种情况下，你需返回 "/" 。
此外，路径中也可能包含多个斜杠 '/' ，如 "/home//foo/" 。
在这种情况下，你可忽略多余的斜杠，返回 "/home/foo" 。
 */

func simplifyPath(path string) string {
	if len(path) <= 1 {
		return path
	}
	start := 0
	buffer := []string{}
	i := 0
	for ; i < len(path); i++ {
		if path[i] == '/' {
			if start < i {
				partPath := path[start:i]
				if partPath == ".." {
					if len(buffer) > 0 {
						buffer = buffer[0:len(buffer)-1]
					}
				} else if partPath != "." {
					buffer = append(buffer, partPath)
				}
			}
			start = i + 1
		}
	}
	if start < len(path) && path[len(path)-1] != '/' {
		partPath := path[start:]
		if partPath == ".." {
			if len(buffer) > 0 {
				buffer = buffer[0:len(buffer)-1]
			}
		} else if partPath != "." {
			buffer = append(buffer, partPath)
		}
	}

	result := "/"
	if len(buffer) > 0 {
		result += buffer[0]
	}
	for i := 1; i < len(buffer); i++ {
		result += "/" + buffer[i]
	}
	return result
}

func TestsimplifyPath() {
	tests := []struct {
		path   string
		result string
	}{
		{
			"/a/./b/../c",
			"/a/c",
		},
		{
			"/a/./b/../../c",
			"/c",
		},
		{
			"/a/./b/../../c/",
			"/c",
		},
		{
			"/a/./b/../../../c/",
			"/c",
		},
		{
			"/a/./b/./././././../../../c/",
			"/c",
		},
		{
			"/home/",
			"/home",
		},
		{
			"/home/",
			"/home",
		},
		{
			"/..",
			"/",
		},
		{
			"/../",
			"/",
		},
		{
			"//home//work//abc///////",
			"/home/work/abc",
		},
		{
			"//a//.//b//..//../c",
			"/c",
		},
	}

	for _, test := range tests {
		fmt.Println(simplifyPath(test.path) == test.result)
	}
}
