package trie

// Node represents a Trie Node
type Node struct {
	Value    string
	Children *NodeMap
	IsEnd    bool
	Word     string
}

// NewNode returns a new node
func NewNode(value string) *Node {
	return &Node{
		Value:    value,
		Children: NewNodeMap(),
	}
}
