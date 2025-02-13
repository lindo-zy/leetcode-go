package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	k := len(word)

	var dfs func(i, j, start int) bool
	dfs = func(i, j, start int) bool {
		tmp := board[i][j]
		board[i][j] = '#'
		if start == k {
			return true
		}
		res := []bool{}
		dis := [][]int{
			{0, 1},
			{1, 0},
			{-1, 0},
			{0, -1},
		}
		for _, d := range dis {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == word[start] {
				res = append(res, dfs(x, y, start+1))
			}
		}
		board[i][j] = tmp
		for _, r := range res {
			if r {
				return true
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 1) {
					return true
				}
			}
		}
	}

	return false

}
