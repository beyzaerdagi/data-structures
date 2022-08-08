package binaryheap

func (heap *MaxBinaryHeap) Insert(val int32) {

	heap.Values = append(heap.Values, val)
	heapify()
}

func heapify() {
	heap := &MaxBinaryHeap{}
	idx := len(heap.Values) - 1
	val := heap.Values[idx]
	for {
		parentIdx := (idx - 1) / 2
		parent := heap.Values[parentIdx]
		if val <= parent {
			break
		}
		heap.Values[parentIdx] = val
		heap.Values[idx] = parent
		idx = parentIdx
	}
}
