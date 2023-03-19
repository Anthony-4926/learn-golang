package Graph

import "math"

func Floyed(G *Graph) (pathes, dis [][]int) {
	n := len(G.Nodes)
	pathes, dis = make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		pathes[i] = make([]int, n)
		dis[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if G.Nodes[i].Edges[j] != nil {
				dis[i][j] = G.Nodes[i].Edges[j].Weight
			} else {
				dis[i][j] = math.MaxInt
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dis[i][k] == math.MaxInt || dis[k][j] == math.MaxInt || i == j {
					continue
				}
				if dis[i][k]+dis[k][j] < dis[i][j] {
					dis[i][j] = dis[i][k] + dis[k][j]
					pathes[i][j] = k
				}
			}
		}
	}
	return

}

// 获得两点之间的路径
func getPath(pathes [][]int, from, to int, path []int) []int {
	k := pathes[from][to]
	if k == -1 {
		return path
	}
	path = getPath(pathes, from, k, path)
	path = append(path, k+1)
	path = getPath(pathes, k, to, path)
	return path
}
