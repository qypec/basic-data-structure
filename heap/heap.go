package heap

import "math"

const MinHeap = 1
const MaxHeap = 2

type Heap struct {
	arr  []element
	size int
	kind int
}

type element struct {
	index    int
	priority int

	Value interface{}
}

func New(heapType int) *Heap { return new(Heap).Init(heapType) }

func NewElement(index, priority int, value interface{}) element {
	return element{index, priority, value}
}

// Builds a min-heap from a priority array.
// The value of the element with priority[i] is equal to values[i]
// Complexity: n
// size of priority and values must be equals
func BuildMinHeap(priority []int, values []interface{}) *Heap {
	var h Heap

	h.Init(MinHeap)
	for i := 0; i < len(priority); i++ {
		h.arr = append(h.arr, element{i + 1, priority[i], values[i]})
		h.size++
	}
	h.heapify()
	return &h
}

// Builds a max-heap from a priority array.
// The value of the element with priority[i] is equal to values[i]
// Complexity: n
// size of priority and values must be equals
func BuildMaxHeap(priority []int, values []interface{}) *Heap {
	var h Heap

	h.Init(MaxHeap)
	for i := 0; i < len(priority); i++ {
		h.arr = append(h.arr, element{i + 1, priority[i], values[i]})
		h.size++
	}
	h.heapify()
	return &h
}

// returns number of elements on the heap
func (h Heap) Size() int { return h.size }

// returns last element on the heap
func (h Heap) Back() *element {
	if h.Size() != 0 {
		return &h.arr[h.Size()]
	}
	return nil
}

// returns first element on the heap
func (h Heap) Front() *element {
	if h.Size() != 0 {
		return &h.arr[1]
	}
	return nil
}

// initializes a heap of size 0.
func (h *Heap) Init(heapType int) *Heap {
	h.arr = make([]element, 1)
	h.arr[0] = element{0, 0, nil}
	h.size = 0
	h.kind = heapType
	return h
}

// adds a new item with priority and value
// complexity: log(n)
func (h *Heap) Insert(priority int, value interface{}) {
	h.arr = append(h.arr, element{h.size + 1, priority, value})
	h.size++
	h.siftingUp(h.size)
}

// returns the value and priority of min element
// removes min element from the heap
// complexity:
// 		(MinHeap): log(n)
//		(MaxHeap): n/2 + log(n)
func (h *Heap) ExtractMin() (interface{}, int) {
	if h.size == 0 {
		return nil, 0
	}
	var min element

	if h.kind == MinHeap {
		frontElem := h.Front()
		backElem := h.Back()
		min = *h.Front()
		h.swap(&frontElem, &backElem)
		h.arr = h.arr[:h.Size()]
		h.size--
		h.siftingDown(1)
	}

	if h.kind == MaxHeap {
		min = *h.Front()
		for i := int(math.Ceil(float64((h.size)/2) + 0.1)); i <= h.size; i++ {
			if h.arr[i].priority < min.priority {
				min = h.arr[i]
			}
		}
		MinElem := &h.arr[min.index]
		backElem := h.Back()
		h.swap(&MinElem, &backElem)
		h.arr = h.arr[:h.Size()]
		h.size--
		h.siftingUp(min.index)
	}

	return min.Value, min.priority
}

// returns the value and priority of max element
// removes max element from the heap
// complexity:
// 		(MaxHeap): log(n)
//		(MinHeap): n/2 + log(n)
func (h *Heap) ExtractMax() (interface{}, int) {
	if h.size == 0 {
		return nil, 0
	}
	var max element

	if h.kind == MaxHeap {
		frontElem := h.Front()
		backElem := h.Back()
		max = *h.Front()
		h.swap(&frontElem, &backElem)
		h.arr = h.arr[:h.Size()]
		h.size--
		h.siftingDown(1)
	}

	if h.kind == MinHeap {
		max = *h.Front()
		for i := int(math.Ceil(float64(h.size / 2))); i <= h.size; i++ {
			if h.arr[i].priority > max.priority {
				max = h.arr[i]
			}
		}
		MaxElem := &h.arr[max.index]
		backElem := h.Back()
		h.swap(&MaxElem, &backElem)
		h.arr = h.arr[:h.Size()]
		h.size--
		h.siftingUp(max.index)
	}

	return max.Value, max.priority
}


