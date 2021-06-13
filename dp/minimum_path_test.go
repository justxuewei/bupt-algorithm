package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumPath(t *testing.T) {
	// case1
	case1 := NewGraph(5)
	case1.AddEdge(0, 1, 2)
	case1.AddEdge(0, 2, 3)
	case1.AddEdge(0, 3, 4)
	case1.AddEdge(1, 2, 4)
	case1.AddEdge(1, 4, 5)
	case1.AddEdge(2, 4, 1)
	case1.AddEdge(3, 4, 3)
	case1rst := minimumPath(case1)
	assert.Equal(t, 4, case1rst)
	// case2
	case2 := NewGraph(5)
	case2.AddEdge(0, 1, 5)
	case2.AddEdge(0, 2, 10)
	case2.AddEdge(0, 3, 4)
	case2.AddEdge(1, 2, 1)
	case2.AddEdge(1, 4, 8)
	case2.AddEdge(2, 4, 2)
	case2.AddEdge(3, 4, 6)
	case2rst := minimumPath(case2)
	assert.Equal(t, 8, case2rst)
}
