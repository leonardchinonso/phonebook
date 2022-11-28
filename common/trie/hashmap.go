package trie

// NodeMap represents a hashset
type NodeMap map[string]*Node

// NewNodeMap returns a new NodeMap pointer
func NewNodeMap() *NodeMap {
	return &NodeMap{}
}

// Add puts a node in NodeMap
func (m *NodeMap) Add(key string, node *Node) {
	(*m)[key] = node
}

// Contains checks if a node is in NodeMap
func (m *NodeMap) Contains(key string) bool {
	_, ok := (*m)[key]
	return ok
}

// Get returns a node from the key
func (m *NodeMap) Get(key string) *Node {
	if !m.Contains(key) {
		panic("node not in the nodemap")
	}
	node, _ := (*m)[key]
	return node
}

// Remove takes out a node from NodeMap
func (m *NodeMap) Remove(key string) {
	delete(*m, key)
}
