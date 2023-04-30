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
	Val     int
	In, Out int
	Nexts   []*Node
	Edges   map[int]*Edge
}

func NewNode(val int) *Node {
	return &Node{
		Val:   val,
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
