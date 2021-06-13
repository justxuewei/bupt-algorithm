package dp

import "math"

func minimumPath(graph Graph) int {
	if len(graph) == 0 && len(graph) != len(graph[0]) {
		return 0
	}
	dp := make([]int, len(graph))
	for i:=0; i<graph.NumOfNodes()-1; i++ {
		dp[i] = math.MaxInt64
	}
	dp[graph.NumOfNodes()-1] = 0
	minimumPathDP(graph, dp, graph.NumOfNodes()-1)
	return dp[0]
}

func minimumPathDP(graph Graph, dp []int, node int) {
	if node == 0 {
		return
	}
	for i:=0; i<len(graph); i++ {
		if graph[i][node] != 0 && dp[i] > dp[node]+graph[i][node] {
			dp[i] = dp[node]+graph[i][node]
			minimumPathDP(graph, dp, i)
		}
	}
}
