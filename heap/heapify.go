package heap

func outOfRange(x, l, r int) bool {
	if x >= l && x <= r {
		return false
	}
	return true
}

// normalizes a heap. Element order respects heap properties
func (h *Heap) heapify() {
	for i := int(h.size / 2); i > 0; i-- {
		h.siftingDown(i)
	}
}

func (h *Heap) siftingSwapCondition(childPriority, parentPriority int) bool {
	if h.kind == MinHeap {
		return childPriority < parentPriority
	} else {
		return childPriority > parentPriority
	}
}

// moves an element up until it satisfies the properties of the heap
func (h *Heap) siftingUp(index int) {
	if outOfRange(index, 1, h.size) {
		return
	}

	child := &h.arr[index]
	if child == nil {
		return
	}
	for parent := h.getParent(child); parent != nil; parent = h.getParent(child) {
		if h.siftingSwapCondition(child.priority, parent.priority) {
			h.swap(&child, &parent)
		} else {
			break
		}
	}
}

// moves an element down until it satisfies the properties of the heap
func (h *Heap) siftingDown(index int) {
	if outOfRange(index, 1, h.size) {
		return
	}

	parent := &h.arr[index]
	if parent == nil {
		return
	}
	for child := h.getChild(parent); child != nil; child = h.getChild(parent) {
		if h.siftingSwapCondition(child.priority, parent.priority) {
			h.swap(&child, &parent)
		} else {
			break
		}
	}
}

// returns the parent of this element
func (h Heap) getParent(child *element) *element {
	parentIndex := int(child.index / 2)
	if parentIndex != 0 {
		return &h.arr[parentIndex]
	}
	return nil
}

// returns the largest(MaxHeap)/less(MinHeap) child of this element
func (h Heap) getChild(parent *element) *element {
	childIndexLeft, childIndexRight := int(parent.index*2), int(parent.index*2+1)
	if outOfRange(childIndexLeft, 1, h.Size()) && outOfRange(childIndexRight, 1, h.Size()) {
		return nil
	} else if outOfRange(childIndexRight, 1, h.Size()) {
		return &h.arr[childIndexLeft]
	} else {
		if h.siftingSwapCondition(h.arr[childIndexLeft].priority, h.arr[childIndexRight].priority) {
			return &h.arr[childIndexLeft]
		} else {
			return &h.arr[childIndexRight]
		}
	}
}

// swap elements inside heap.arr
// heap.index remains untouched
func (h *Heap) swap(child, parent **element) {
	childIndex := (*child).index
	parentIndex := (*parent).index

	h.arr[childIndex], h.arr[parentIndex] = h.arr[parentIndex], h.arr[childIndex]
	h.arr[parentIndex].index, h.arr[childIndex].index = parentIndex, childIndex

	*child = &h.arr[parentIndex]
	*parent = &h.arr[childIndex]
}
