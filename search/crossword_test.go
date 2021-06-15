package search

import "testing"

func TestFindWords(t *testing.T) {
	case1Boaad := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	case1Words := []string{"oath", "pea", "eat", "rain"}
	case1Ret := findWords(case1Boaad, case1Words)
	t.Log(case1Ret)
	case2Board := [][]byte{
		{'a', 'b'},
	}
	case2Words := []string{"ba"}
	case2Ret := findWords(case2Board, case2Words)
	t.Log(case2Ret)
}
