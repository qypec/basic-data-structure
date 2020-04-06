package heap

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSwap(t *testing.T) {
	expectedResult := &Heap{
		[]element{
			{0, 0, nil},
			{1, 20, "second"},
			{2, 10, "first"},
		},
		2,
		MaxHeap,
	}
	heap := &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
		},
		2,
		MaxHeap,
	}

	first := &heap.arr[1]
	second := &heap.arr[2]
	heap.swap(&first, &second)
	require.Equal(t, expectedResult, heap)
}

func TestGetChild(t *testing.T) {
	/* Test01 one child */
	heap := &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
		},
		2,
		MaxHeap,
	}

	actual := heap.getChild(&heap.arr[1])
	require.Equal(t, &element{2, 20, "second"}, actual)

	/* Test02 no child */
	heap = &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
		},
		2,
		MaxHeap,
	}

	if actual = heap.getChild(&heap.arr[2]); actual != nil {
		t.Error(fmt.Sprintf("expected: %v\n actual: %v\n", nil, actual))
	}

	/* Test03 larger child */
	heap = &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
			{3, 30, "third"},
		},
		3,
		MaxHeap,
	}

	actual = heap.getChild(&heap.arr[1])
	require.Equal(t, &element{3, 30, "third"}, actual)

	/* Test04 less child */
	heap = &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
			{3, 30, "third"},
		},
		3,
		MinHeap,
	}

	actual = heap.getChild(&heap.arr[1])
	require.Equal(t, &element{2, 20, "second"}, actual)
}

func TestGetParent(t *testing.T) {
	/* Test01 */
	heap := &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
		},
		2,
		MaxHeap,
	}

	actual := heap.getParent(&heap.arr[2])
	require.Equal(t, &element{1, 10, "first"}, actual)

	/* Test02 no parent */
	heap = &Heap{
		[]element{
			{0, 0, nil},
			{1, 10, "first"},
			{2, 20, "second"},
		},
		2,
		MaxHeap,
	}

	if actual = heap.getParent(&heap.arr[1]); actual != nil {
		t.Error(fmt.Sprintf("\nexpected: %v\n actual: %v\n", nil, actual))
	}
}

type myTestSifting struct {
	h           Heap
	siftingElem int
}

func TestSiftingUp(t *testing.T) {
	expectedResult := []Heap{
		/* 1 */ { // no changes (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
			},
			2,
			MinHeap,
		},
		/* 2 */ { // no changes (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
			},
			2,
			MaxHeap,
		},
		/* 3 */ { // last elem sifting up (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 5, "new first"},
				{2, 20, "second"},
				{3, 10, "first"},
			},
			3,
			MinHeap,
		},
		/* 4 */ { // last elem sifting up (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 50, "new first"},
				{2, 10, "second"},
				{3, 20, "first"},
			},
			3,
			MaxHeap,
		},
		/* 5 */ { // first elem sifting up (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
				{3, 30, "third"},
			},
			3,
			MinHeap,
		},
		/* 6 */ { // first elem sifting up (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
				{3, 5, "third"},
			},
			3,
			MaxHeap,
		},
		/* 7 */ { // mid elem sifting up (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 5, "mid"},
				{2, 10, "first"},
				{3, 30, "third"},
			},
			3,
			MinHeap,
		},
		/* 8 */ { // mid elem sifting up (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 30, "mid"},
				{2, 20, "first"},
				{3, 5, "third"},
			},
			3,
			MaxHeap,
		},
	}

	actual := []myTestSifting{
		/* 1 */ {Heap{ // no changes (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
			},
			2,
			MinHeap,
		},
			1,
		},
		/* 2 */ {Heap{ // no changes (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
			},
			2,
			MaxHeap,
		},
			1,
		},
		/* 3 */ {Heap{ // last elem sifting up (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
				{3, 5, "new first"},
			},
			3,
			MinHeap,
		},
			3,
		},
		/* 4 */ {Heap{ // last elem sifting up (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
				{3, 50, "new first"},
			},
			3,
			MaxHeap,
		},
			3,
		},
		/* 5 */ {Heap{ // first elem sifting up (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
				{3, 30, "third"},
			},
			3,
			MinHeap,
		},
			1,
		},
		/* 6 */ {Heap{ // first elem sifting up (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
				{3, 5, "third"},
			},
			3,
			MaxHeap,
		},
			1,
		},
		/* 7 */ {Heap{ // mid elem sifting up (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 5, "mid"},
				{3, 30, "third"},
			},
			3,
			MinHeap,
		},
			2,
		},
		/* 8 */ {Heap{ // mid elem sifting up (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 30, "mid"},
				{3, 5, "third"},
			},
			3,
			MaxHeap,
		},
			2,
		},
	}

	for i := 0; i < len(actual); i++ {
		testCase := actual[i]
		testCase.h.siftingUp(testCase.siftingElem)
		require.Equal(t, expectedResult[i], testCase.h, fmt.Sprintf("testCase: %v", i+1))
	}
}

func TestSiftingDown(t *testing.T) {
	expectedResult := []Heap{
		/* 1 */ { // no changes (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
			},
			2,
			MinHeap,
		},
		/* 2 */ { // no changes (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
			},
			2,
			MaxHeap,
		},
		/* 3 */ { // first elem sifting down to mid (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 5, "second"},
				{2, 10, "new mid"},
				{3, 20, "last"},
			},
			3,
			MinHeap,
		},
		/* 4 */ { // first elem sifting down to mid (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 50, "new first"},
				{2, 10, "second"},
				{3, 30, "new mid"},
			},
			3,
			MaxHeap,
		},
		/* 5 */ { // last elem sifting down (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
				{3, 30, "third"},
			},
			3,
			MinHeap,
		},
		/* 6 */ { // last elem sifting down (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
				{3, 5, "third"},
			},
			3,
			MaxHeap,
		},
	}

	actual := []myTestSifting{
		/* 1 */ {Heap{ // no changes (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
			},
			2,
			MinHeap,
		},
			1,
		},
		/* 2 */ {Heap{ // no changes (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
			},
			2,
			MaxHeap,
		},
			1,
		},
		/* 3 */ {Heap{ // first elem sifting down to mid (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "new mid"},
				{2, 5, "second"},
				{3, 20, "last"},
			},
			3,
			MinHeap,
		},
			1,
		},
		/* 4 */ {Heap{ // first elem sifting down to mid (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 30, "new mid"},
				{2, 10, "second"},
				{3, 50, "new first"},
			},
			3,
			MaxHeap,
		},
			1,
		},
		/* 5 */ {Heap{ // last elem sifting down (MinHeap)
			[]element{
				{0, 0, nil},
				{1, 10, "first"},
				{2, 20, "second"},
				{3, 30, "third"},
			},
			3,
			MinHeap,
		},
			3,
		},
		/* 6 */ {Heap{ // last elem sifting down (MaxHeap)
			[]element{
				{0, 0, nil},
				{1, 20, "first"},
				{2, 10, "second"},
				{3, 5, "third"},
			},
			3,
			MaxHeap,
		},
			3,
		},
	}

	for i := 0; i < len(actual); i++ {
		testCase := actual[i]
		testCase.h.siftingDown(testCase.siftingElem)
		require.Equal(t, expectedResult[i], testCase.h, fmt.Sprintf("testCase: %v", i+1))
	}
}

func TestHeapify(t *testing.T) {
	expectedResult := []Heap{
		/* 1 */ {
			[]element{
				{0, 0, nil},
				{1, 2, nil},
				{2, 3, nil},
				{3, 5, nil},
				{4, 4, nil},
				{5, 6, nil},
				{6, 7, nil},
			},
			6,
			MinHeap,
		},
		/* 2 */ {
			[]element{
				{0, 0, nil},
				{1, 7, nil},
				{2, 6, nil},
				{3, 4, nil},
				{4, 5, nil},
				{5, 2, nil},
				{6, 3, nil},
			},
			6,
			MaxHeap,
		},
	}

	actual := []Heap{
		/* 1 */ {
			[]element{
				{0, 0, nil},
				{1, 7, nil},
				{2, 6, nil},
				{3, 5, nil},
				{4, 4, nil},
				{5, 3, nil},
				{6, 2, nil},
			},
			6,
			MinHeap,
		},
		/* 2 */ {
			[]element{
				{0, 0, nil},
				{1, 2, nil},
				{2, 7, nil},
				{3, 4, nil},
				{4, 5, nil},
				{5, 6, nil},
				{6, 3, nil},
			},
			6,
			MaxHeap,
		},
	}

	for i := 0; i < len(actual); i++ {
		testCase := actual[i]
		testCase.heapify()
		require.Equal(t, expectedResult[i], testCase, fmt.Sprintf("testCase: %v", i+1))
	}
}
