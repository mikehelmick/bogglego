// Package trie implements a basic string based prefix tree.
package trie

import (
	"fmt"
	"strings"
)

// Error in manipulation of a Trie
type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

type node struct {
	value    string
	valid    bool
	children map[string]*node
}

// Trie is the representation of the prefix tree.
type Trie struct {
	roots map[string]*node
}

// New initializes a new, empty Trie.
func New() *Trie {
	return &Trie{make(map[string]*node)}
}

// insert is a recursive method to add strings to the dictionary, character by
// character.
func insert(n *node, s string) {
	if len(s) == 0 {
		// end of the insert chain, this is a valid word.
		n.valid = true
		return
	}
	f := s[0:1]
	l := s[1:]

	if _, ok := n.children[f]; !ok {
		n.children[f] = &node{f, false, make(map[string]*node)}
	}
	insert(n.children[f], l)
}

// AddWord ensures that w is added to the dictionary.
func (t *Trie) AddWord(w string) error {
	if len(w) == 0 {
		return &Error{"Word to add can't be empty."}
	}
	w = strings.ToLower(w)
	for i := 0; i < len(w); i++ {
		ch := w[i]
		if ch < 'a' || ch > 'z' {
			return &Error{fmt.Sprintf("Invalid char '%v'", ch)}
		}
	}

	first := w[0:1]
	rest := w[1:]

	if _, ok := t.roots[first]; !ok {
		t.roots[first] = &node{first, len(rest) == 0, make(map[string]*node)}
	}
	insert(t.roots[first], rest)

	return nil
}

func search(t *Trie, s string) *node {
	if len(s) == 0 {
		return nil
	}
	first := s[0:1]
	rest := s[1:]

	return crawl(t.roots[first], rest)
}

func crawl(n *node, s string) *node {
	if n == nil {
		return nil
	}
	if len(s) == 0 {
		return n
	}
	first := s[0:1]
	rest := s[1:]

	return crawl(n.children[first], rest)
}

// IsPrefix returns true if a given string is a valid prefix in this Trie.
func (t *Trie) IsPrefix(s string) bool {
	node := search(t, s)
	if node == nil {
		return false
	}
	return len(node.children) > 0
}

// IsWord returns true if a given string is a valid word in the Trie
func (t *Trie) IsWord(s string) bool {
	node := search(t, s)
	if node == nil {
		return false
	}
	return node.valid
}
