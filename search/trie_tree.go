package search

type trieNode struct {
	children map[byte]*trieNode
	word     string
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[byte]*trieNode),
	}
}

func buildTrieTree(words []string) *trieNode {
	root := newTrieNode()
	var lastNode *trieNode
	for _, word := range words {
		lastNode = root
		for i := range word {
			node, ok := lastNode.children[word[i]]
			if !ok {
				node = newTrieNode()
				lastNode.children[word[i]] = node
			}
			lastNode = node
		}
		lastNode.word = word
	}
	return root
}
