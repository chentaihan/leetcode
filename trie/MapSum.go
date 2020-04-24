package trie

import "fmt"

type TrieNode1 struct {
	Count    int
	IsEnd    bool
	Children [26]*TrieNode1
}

type MapSum struct {
	root *TrieNode1
}

/** Initialize your data structure here. */
func Constructor1() MapSum {
	return MapSum{root: &TrieNode1{}}
}

func (this *MapSum) Insert(word string, val int) {
	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.Children[index] == nil {
			node.Children[index] = &TrieNode1{
				Count:    0,
				Children: [26]*TrieNode1{},
			}
		}
		node = node.Children[index]
	}
	node.Count = val
	node.IsEnd = true
}

func (this *MapSum) Sum(word string) int {
	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.Children[index] == nil {
			return 0
		}
		node = node.Children[index]
	}
	return this.sum(node)
}

func (this *MapSum) sum(node *TrieNode1) int {
	result := 0
	if node.IsEnd {
		result += node.Count
	}
	for i := 0; i < 26; i++ {
		if node.Children[i] != nil {
			result += this.sum(node.Children[i])
		}
	}
	return result
}

func TestMapSum() {
	tests := []struct {
		word string
		val  int
	}{
		{
			"apple",
			3,
		},
		{
			"app",
			2,
		},
	}
	sm := Constructor1()
	for _, test := range tests {
		sm.Insert(test.word, test.val)
	}
	fmt.Println(sm.Sum("ap"))
}
