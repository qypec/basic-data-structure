package hash_table

import "container/list"

type HashTable struct {
	arr []*list.List
	size int // number of elements
}


func polynomialHashFunction(str string) int {

}

func (h *HashTable) Init() {
	h.arr = make([]*list.List, 20)
	h.size = 0
}

func (h *HashTable) Add(str string) {
	hash := polynomialHashFunction(str)
	h.arr[hash].PushFront(str)
	h.size++
}

func (h *HashTable) Delete(str string) {
	hash := polynomialHashFunction(str)
	for e := h.arr[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == str {
			h.arr[hash].Remove(e)
			h.size--
		}
	}
}

func (h *HashTable) Find(str string) bool {
	hash := polynomialHashFunction(str)
	for e := h.arr[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == str {
			return true
		}
	}
	return false
}
