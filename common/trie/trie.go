package trie

import "fmt"

// Trie represents the trie data structure
type Trie struct {
	Root *Node
}

// NewTrie returns a Trie pointer
func NewTrie() *Trie {
	root := NewNode("")
	return &Trie{
		Root: root,
	}
}

// AddWord adds a new word to the trie
func (t *Trie) AddWord(word string) (*Node, error) {
	if word == "" {
		return nil, fmt.Errorf("cannot add empty word to trie")
	}

	curr := t.Root            // copy the root node to a temp variable
	for i, ch := range word { // iterate through the word and search on each character
		chStr := string(ch)
		var node *Node

		// if the current character is in this trie level, move down
		if curr.Children.Contains(chStr) {
			// set the curr variable to the next
			node = curr.Children.Get(chStr)
		} else {
			// if the current character is not in this trie level, add it
			node = NewNode(chStr)
			curr.Children.Add(chStr, node) // add the current node as a child
		}

		curr = node

		// if the character is the last character in the word, set IsEnd
		if i == len(word)-1 {
			curr.IsEnd = true
			curr.Word = word
		}
	}

	return curr, nil
}

// Find looks for a substring in the trie that matches word and returns the last node
func (t *Trie) Find(word string) (*Node, error) {
	if word == "" {
		return nil, fmt.Errorf("cannot search empty word")
	}

	curr := t.Root // copy the root node to a temp variable
	for _, ch := range word {
		chStr := string(ch)

		if !curr.Children.Contains(chStr) {
			return nil, fmt.Errorf("word not in trie")
		}

		curr = curr.Children.Get(chStr)
	}

	return curr, nil
}

// FindWord looks for the exact word match in the trie and returns the last node
func (t *Trie) FindWord(word string) (*Node, error) {
	curr, err := t.Find(word)
	if err != nil {
		return nil, err
	}

	// if the word is not a registered substring that was added, return error
	if !curr.IsEnd {
		return nil, fmt.Errorf("word not in trie")
	}

	return curr, nil
}

// DeleteWord sets the word last character to false
// so that the word is no longer reachable by FindWord
func (t *Trie) DeleteWord(node *Node) error {
	if !node.IsEnd {
		return fmt.Errorf("cannot delete a node if it is not a word")
	}
	node.IsEnd = false
	node.Word = ""
	return nil
}

// UpdateWord calls DeleteWord and then AddWord
func (t *Trie) UpdateWord(node *Node, newWord string) (*Node, error) {
	if err := t.DeleteWord(node); err != nil {
		return nil, err
	}
	node, err := t.AddWord(newWord)
	if err != nil {
		return nil, err
	}
	return node, nil
}
