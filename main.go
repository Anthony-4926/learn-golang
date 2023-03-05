package main

import (
	"fmt"
)

func main() {
	M1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	M2 := [][]int{{0, 1}, {0, 0}}
	fmt.Println(process(M1))
	fmt.Println(process(M2))
}

func process(M [][]int) int {
	dp := make([][]int, len(M)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(M[0])+1)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	return dfs(M, 0, 0, dp)
}

func dfs(M [][]int, x, y int, dp [][]int) int {

	if x == len(M)-1 && y == len(M[0])-1 {
		dp[x][y] = 1
		return 1
	}
	if dp[x][y] != -1 {
		return dp[x][y]
	}
	if x > len(M)-1 {
		dp[x][y] = 0
		return 0
	}
	if y > len(M[0])-1 {
		dp[x][y] = 0
		return 0
	}
	if M[x][y] == 1 {
		dp[x][y] = 0
		return 0
	}

	p1 := dfs(M, x+1, y, dp)
	p2 := dfs(M, x, y+1, dp)
	dp[x][y] = p1 + p2
	return dp[x][y]
}
