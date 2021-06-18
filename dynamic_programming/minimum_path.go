package dynamic_programming

import "algorithm/foundation"

func minimumPath(graph Graph) int {
	if len(graph) == 0 && len(graph) != len(graph[0]) {
		return 0
	}
	dp := make([]int, len(graph))
	for i := 1; i < graph.NumOfNodes(); i++ {
		dp[i] = Inf
	}
	dp[0] = 0
	getMinimumPath(graph, dp, graph.NumOfNodes()-1)
	return dp[graph.NumOfNodes()-1]
}

func getMinimumPath(graph Graph, dp []int, node int) {
	if dp[node] != Inf {
		return
	}
	minpath := Inf
	for i := 0; i < graph.NumOfNodes(); i++ {
		if graph[i][node] != Inf {
			if dp[i] == Inf {
				getMinimumPath(graph, dp, i)
			}
			minpath = foundation.Min(minpath, dp[i]+graph[i][node])
		}
	}
	dp[node] = minpath
}
