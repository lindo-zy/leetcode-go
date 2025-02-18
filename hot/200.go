package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = '0'
		dst := [][]int{
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		}
		for _, d := range dst {
			nx, ny := x+d[0], y+d[1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == '1' {
				dfs(nx, ny)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans += 1
			}
		}
	}
	return ans
}
