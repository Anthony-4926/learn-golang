package Graph

import "math"

func Dijkstra(G *Graph, from *Node) map[*Node]int {
	distanceMap := map[*Node]int{}
	distanceMap[from] = 0

	selectedNodes := map[*Node]struct{}{}
	minNode := getMinDistanceAndUnselectedNode(distanceMap, selectedNodes)
	for minNode != nil {
		distance := distanceMap[minNode]
		for _, edge := range minNode.Edges {
			toNode := G.Nodes[edge.To]
			if _, ok := distanceMap[toNode]; !ok {
				distanceMap[toNode] = distance + edge.Weight
			} else {
				distanceMap[toNode] = min(distanceMap[toNode], distance+edge.Weight)
			}
		}
		selectedNodes[minNode] = struct{}{}
	}
	return distanceMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, touchedNodes map[*Node]struct{}) *Node {
	var minNode *Node = nil
	minDistance := math.MaxInt
	for node, d := range distanceMap {
		if _, ok := touchedNodes[node]; !ok && d < minDistance {
			minNode = node
			minDistance = d
		}
	}
	return minNode
}
