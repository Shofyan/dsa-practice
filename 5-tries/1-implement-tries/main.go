package main

import (
	"fmt"
)

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// NewTrie initializes and returns a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith checks if any word in the Trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Display prints the entire Trie structure
func (t *Trie) Display() {
	fmt.Println("Trie structure:")
	t.displayHelper(t.root, "", 0)
}

// displayHelper is a recursive helper function to display the Trie
func (t *Trie) displayHelper(node *TrieNode, prefix string, depth int) {
	if node.isEnd {
		fmt.Printf("%sWord: %s\n", getIndent(depth), prefix)
	}

	for ch, childNode := range node.children {
		fmt.Printf("%s'%c'\n", getIndent(depth), ch)
		t.displayHelper(childNode, prefix+string(ch), depth+1)
	}
}

// getIndent returns indentation string for tree visualization
func getIndent(depth int) string {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	return indent
}

func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "banana", "band", "bandana"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Display the Trie structure
	trie.Display()
	fmt.Println()

	// Test Search
	fmt.Println("Search 'apple':", trie.Search("apple"))     // true
	fmt.Println("Search 'app':", trie.Search("app"))         // true
	fmt.Println("Search 'appl':", trie.Search("appl"))       // false
	fmt.Println("Search 'banana':", trie.Search("banana"))   // true
	fmt.Println("Search 'band':", trie.Search("band"))       // true
	fmt.Println("Search 'bandana':", trie.Search("bandana")) // true
	fmt.Println("Search 'ban':", trie.Search("ban"))         // false

	// Test StartsWith
	fmt.Println("StartsWith 'app':", trie.StartsWith("app")) // true
	fmt.Println("StartsWith 'ban':", trie.StartsWith("ban")) // true
	fmt.Println("StartsWith 'cat':", trie.StartsWith("cat")) // false
}
