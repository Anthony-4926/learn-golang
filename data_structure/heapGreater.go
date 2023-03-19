package data_structure

type HeapNode map[interface{}]interface{}

type HeapGreater struct {
	heap         []interface{}
	inverseIndex map[interface{}]int
	heapSize     int
	less         func(a, b interface{}) bool
}

func NewHeapGreater(less func(a, b interface{}) bool) *HeapGreater {
	return &HeapGreater{
		less:         less,
		inverseIndex: map[interface{}]int{},
		heap:         []interface{}{},
	}
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
	h.Remove(h.heap[0])
}

func (h *HeapGreater) Remove(obj interface{}) {
	index := h.inverseIndex[obj]
	h.swap(index, h.heapSize-1)
	delete(h.inverseIndex, obj)
	h.heapSize--
	h.heap = h.heap[:h.heapSize]
	if index < h.heapSize {
		h.resign(index)
	}

}

func (h *HeapGreater) resign(index int) {
	h.heapInsert(index)
	h.heapify(index)
}

func (h *HeapGreater) Add(obj interface{}) {
	h.heap = append(h.heap, obj)
	h.heapSize++
	h.inverseIndex[obj] = h.heapSize - 1
	h.heapInsert(h.heapSize - 1)

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
		minIndex := index
		if h.less(h.heap[left], h.heap[minIndex]) {
			minIndex = left
		}
		if left+1 < h.heapSize && h.less(h.heap[left+1], h.heap[minIndex]) {
			minIndex = left + 1
		}
		if index == minIndex {
			break
		}
		h.swap(minIndex, index)
		left = index*2 + 1
	}
}

func (h *HeapGreater) swap(i int, j int) {
	a, b := h.heap[i], h.heap[j]
	h.inverseIndex[a] = j
	h.inverseIndex[b] = i

	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *HeapGreater) AllElem() []interface{} {
	return h.heap
}
