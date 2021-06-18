package dynamic_programming

import "math"

const Inf = math.MaxInt64

type Graph [][]int

func NewGraph(nNodes int) Graph {
	if nNodes == 0 {
		return nil
	}
	graph := make([][]int, nNodes)
	for i := 0; i < nNodes; i++ {
		graph[i] = make([]int, nNodes)
		for j := 0; j < nNodes; j++ {
			graph[i][j] = Inf
		}
	}
	return graph
}

func (g Graph) AddEdge(from, to, val int) {
	if from >= len(g) || to >= len(g[0]) {
		panic("edge couldn't be added")
	}
	g[from][to] = val
}

func (g Graph) NumOfNodes() int {
	return len(g)
}
