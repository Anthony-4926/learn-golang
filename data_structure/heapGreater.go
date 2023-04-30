package data_structure

type HeapNode struct {
	val int
}

type HeapGreater struct {
	heap     []*HeapNode
	indexMap map[*HeapNode]int
	heapSize int
	less     func(*HeapNode, *HeapNode) bool
}

func NewHeapGreater(c func(*HeapNode, *HeapNode) bool) *HeapGreater {
	return &HeapGreater{
		heap:     make([]*HeapNode, 0),
		indexMap: make(map[*HeapNode]int),
		heapSize: 0,
		less:     c,
	}
}

func (h *HeapGreater) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *HeapGreater) Size() int {
	return h.heapSize
}

func (h *HeapGreater) Contains(obj *HeapNode) bool {
	_, ok := h.indexMap[obj]
	return ok
}

func (h *HeapGreater) Peek() *HeapNode {
	return h.heap[0]
}

func (h *HeapGreater) Push(obj *HeapNode) {
	h.heap = append(h.heap, obj)
	h.indexMap[obj] = h.heapSize
	h.heapInsert(h.heapSize)
	h.heapSize++
}

func (h *HeapGreater) Pop() *HeapNode {
	ans := h.heap[0]
	h.swap(0, h.heapSize-1)
	delete(h.indexMap, ans)
	h.heap = h.heap[:h.heapSize-1]
	h.heapSize--
	h.heapify(0)
	return ans
}

func (h *HeapGreater) Remove(obj *HeapNode) {
	replace := h.heap[h.heapSize-1]
	index := h.indexMap[obj]
	delete(h.indexMap, obj)
	h.heap = h.heap[:h.heapSize-1]
	h.heapSize--
	if obj != replace {
		h.heap[index] = replace
		h.indexMap[replace] = index
		h.Resign(replace)
	}
}

func (h *HeapGreater) Resign(obj *HeapNode) {
	h.heapInsert(h.indexMap[obj])
	h.heapify(h.indexMap[obj])
}

func (h *HeapGreater) GetAllElements() []*HeapNode {
	return h.heap
}

func (h *HeapGreater) heapInsert(index int) {
	for h.less(h.heap[index], h.heap[(index-1)/2]) {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *HeapGreater) heapify(index int) {
	left := index*2 + 1
	for left < h.heapSize {
		best := left + 1
		if best < h.heapSize && h.less(h.heap[best], h.heap[left]) {
			best = left + 1
		} else {
			best = left
		}
		if !h.less(h.heap[best], h.heap[index]) {
			break
		}
		h.swap(best, index)
		index = best
		left = index*2 + 1
	}
}

func (h *HeapGreater) swap(i, j int) {
	o1 := h.heap[i]
	o2 := h.heap[j]
	h.heap[i] = o2
	h.heap[j] = o1
	h.indexMap[o2] = i
	h.indexMap[o1] = j
}
