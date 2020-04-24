package trie

/**
208. 实现 Trie (前缀树)
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

const CHILD_SIZE = 26

type TrieNode struct {
	Val      byte
	IsEnd    bool
	Children [CHILD_SIZE]*TrieNode
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.Children[index] == nil {
			node.Children[index] = &TrieNode{
				Val:      word[i],
				IsEnd:    false,
				Children: [26]*TrieNode{},
			}
		}
		node = node.Children[index]
	}
	node.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return node.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return true
}
