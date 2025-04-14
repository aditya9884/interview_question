/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
(using trie algorithm)
*/

package main

import (
	"fmt"
)

// trie represent each node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	count    int
	isEnd    bool
}

func newNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// trie struct with insert and prefix logic
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = newNode()
		}
		node = node.children[char]
		node.count++
	}
	node.isEnd = true
}

func (t *Trie) longestCommonPrefix(n int) string {
	prefix := ""
	node := t.root

	for {
		if len(node.children) != 1 || node.isEnd {
			break
		}
		for ch, next := range node.children {
			if next.count == n {
				prefix += string(ch)
				node = next
			} else {
				return prefix
			}
		}
	}
	return prefix

}

func longestCommonPrefixTrie(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	trie := NewTrie()
	for _, word := range strs {
		trie.Insert(word)
	}
	return trie.longestCommonPrefix(len(strs))
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefixTrie(strs))
}
