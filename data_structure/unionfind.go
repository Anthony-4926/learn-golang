package data_structure

type node struct {
	value interface{}
}

type UnionFind struct {
	nodes   map[interface{}]node
	parents map[node]node
	sizeMap map[node]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		nodes:   make(map[interface{}]node),
		parents: make(map[node]node),
		sizeMap: make(map[node]int),
	}
}

func NewUnionFindWithValues(values []int) *UnionFind {
	u := &UnionFind{}
	for _, v := range values {
		u.Add(v)
	}
	return u
}

// Add 向并查集中添加一个节点
func (u *UnionFind) Add(v interface{}) {
	if _, ok := u.nodes[v]; !ok {
		vNode := node{v}
		u.nodes[v] = vNode
		u.parents[vNode] = vNode
		u.sizeMap[vNode] = 1
	}
}

// findFather 查找一个集合的代表节点
func (u *UnionFind) findFather(cur node) node {
	path := make([]node, 0)
	for u.parents[cur] != cur {
		path = append(path, cur)
		cur = u.parents[cur]
	}
	for _, p := range path {
		u.parents[p] = cur
	}
	return cur
}

// Union 合并两个集合
func (u *UnionFind) Union(a, b interface{}) {
	bigger := u.findFather(u.nodes[a])
	smaller := u.findFather(u.nodes[b])

	if bigger != smaller {
		if u.sizeMap[smaller] > u.sizeMap[bigger] {
			bigger, smaller = smaller, bigger
		}
		u.parents[smaller] = bigger
		u.sizeMap[bigger] += u.sizeMap[smaller]
		delete(u.sizeMap, smaller)
	}
}

func (u *UnionFind) IsSameSet(a, b interface{}) bool {
	aFather := u.findFather(u.nodes[a])
	bFather := u.findFather(u.nodes[b])
	return aFather == bFather
}

func (u *UnionFind) Size() int {
	return len(u.sizeMap)
}
