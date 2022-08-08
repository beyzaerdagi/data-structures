package binaryheap

func (heap *MaxBinaryHeap) Remove() int32 {

	max := heap.Values[0]
	end := heap.Values[len(heap.Values)-1]
	heap.Values = heap.Values[:len(heap.Values)-1]
	if len(heap.Values) > 0 {
		heap.Values[0] = end
		sinkDown()
	}
	return max
}

func sinkDown() {
	heap := &MaxBinaryHeap{}
	idx := 0
	length := len(heap.Values)
	element := heap.Values[0]

	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2
		leftChild, rightChild := 0, 0
		swap := -1

		if leftChildIdx < length {
			leftChild = int(heap.Values[leftChildIdx])
			if leftChild > int(element) {
				swap = leftChildIdx
			}
		}

		if rightChildIdx < length {
			rightChild = int(heap.Values[rightChildIdx])
			if (swap == -1 && rightChild > int(element)) || (swap != -1 && rightChild > leftChild) {
				swap = rightChildIdx
			}
		}

		if swap == -1 {
			break
		}
		heap.Values[idx] = heap.Values[swap]
		heap.Values[swap] = element
		idx = swap
	}
}
