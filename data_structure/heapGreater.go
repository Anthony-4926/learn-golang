package data_structure

type HeapGreater struct {
	heap         []interface{}
	inverseIndex map[interface{}]int
	heapSize     int
	less         func(a, b interface{}) bool
}

func NewHeapGreater(less func(a, b interface{}) bool) *HeapGreater {
	return &HeapGreater{less: less}
}

func (h *HeapGreater) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *HeapGreater) Size() int {
	return h.heapSize
}

func (h *HeapGreater) Contains(obj interface{}) bool {
	if _, ok := h.inverseIndex[obj]; ok {
		return true
	}
	return false
}

func (h *HeapGreater) Peek() interface{} {
	return h.heap[0]
}

func (h *HeapGreater) Pop() {

}

func (h *HeapGreater) Remove(obj interface{}) {

}

func (h *HeapGreater) Add(obj interface{}) {
	h.heap = append(h.heap, obj)
	h.inverseIndex[obj] = h.heapSize
	h.heapInsert(h.heapSize)
	h.heapSize++
}

func (h *HeapGreater) heapInsert(index int) {
	for h.less(h.heap[index], h.heap[(index-1)/2]) {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *HeapGreater) swap(i int, j int) {
	a, b := h.heap[i], h.heap[j]
	h.inverseIndex[a] = j
	h.inverseIndex[b] = i

	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]

}


