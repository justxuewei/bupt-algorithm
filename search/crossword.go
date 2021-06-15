package search

var directions = [][]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return nil
	}
	// build trie tree
	tireTreeRoot := buildTrieTree(words)
	// visited
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	// auxiliary data
	var ret []string
	var track []byte
	// back tracing
	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[0]); j++ {
			findWordsBackTracing(board, visited, i, j, tireTreeRoot, &track, &ret)
		}
	}
	return ret
}

func findWordsBackTracing(board [][]byte, visited [][]bool, i, j int, parent *trieNode, track *[]byte, ret *[]string) {
	child, ok := parent.children[board[i][j]]
	if !ok {
		return
	}
	visited[i][j] = true
	*track = append(*track, board[i][j])
	if child.word == string(*track) {
		*ret = append(*ret, string(*track))
		// avoid add a word repeatedly
		child.word = ""
	}
	for _, d := range directions {
		if i+d[0] >= 0 && j+d[1] >= 0 && i+d[0] < len(board) && j+d[1] < len(board[0]) && !visited[i+d[0]][j+d[1]] {
			findWordsBackTracing(board, visited, i+d[0], j+d[1], child, track, ret)
		}
	}
	visited[i][j] = false
	*track = (*track)[:len(*track)-1]
}
