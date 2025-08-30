package practice

import (
	"strings"
)

type TrieNode struct {
	data       rune
	child      [26]*TrieNode
	isTerminal bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{data: 0}}
}

func (t *Trie) Insert(word string) {
	temp := t.root
	for _, ch := range word {
		index := ch - 'a'
		if temp.child[index] == nil {
			temp.child[index] = &TrieNode{data: ch}
		}
		temp = temp.child[index]
	}
	temp.isTerminal = true
}

func (t *Trie) FindPrefix(word string) string {
	temp := t.root
	for i, ch := range word {
		index := ch - 'a'
		if temp.child[index] == nil {
			return ""
		}
		temp = temp.child[index]
		if temp.isTerminal {
			return word[:i+1]
		}
	}
	return ""
}

func replaceWords(dictionary []string, sentence string) string {

	res := strings.Builder{}
	trie := NewTrie()

	for _, word := range dictionary {
		trie.Insert(word)
	}

	words := strings.Split(sentence, " ")
	for _, word := range words {
		prefix := trie.FindPrefix(word)
		if prefix == "" {
			res.WriteString(word)
		} else {
			res.WriteString(prefix)
		}
		res.WriteString(" ")
	}

	return strings.TrimSpace(res.String())
}
