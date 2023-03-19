package Graph

type Graph struct {
	Nodes map[int]*Node
	Edges map[*Edge]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: map[int]*Node{},
		Edges: map[*Edge]struct{}{},
	}
}

type Node struct {
	In, Out int
	Nexts   []*Node
	Edges   map[int]*Edge
}

func NewNode() *Node {
	return &Node{
		Nexts: []*Node{},
		Edges: map[int]*Edge{},
	}
}

type Edge struct {
	From, To int
	Weight   int
}

func NewEdge(from int, to int, weight int) *Edge {
	return &Edge{From: from, To: to, Weight: weight}
}
