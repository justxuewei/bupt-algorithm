package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrieTree(t *testing.T) {
	case1 := []string{"oath", "pea", "eat", "rain"}
	case1TrieTreeRoot := buildTrieTree(case1)
	for _, word := range case1 {
		validTrieTree(t, case1TrieTreeRoot, word)
	}
}

func validTrieTree(t *testing.T, root *trieNode, word string) {
	cntNode := root
	for i := range word {
		node, ok := cntNode.children[word[i]]
		assert.True(t, ok)
		cntNode = node
	}
	assert.Equal(t, word, cntNode.word)
}
