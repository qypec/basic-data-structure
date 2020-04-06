package heap

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	var h Heap

	h.Init(MinHeap)
	require.Equal(t, h.size, 0)
	require.Equal(t, len(h.arr), 1)
	require.Equal(t, h.arr[0], element{0, 0, nil})
}

type myTestBuildHeap struct {
	priority []int
	values   []interface{}
}

func TestBuildMaxHeap(t *testing.T) {
	/* Test 01 right format */
	expectedTest01 := &Heap{
		[]element{
			{0, 0, nil},
			{1, 5, "5"},
			{2, 4, "4"},
			{3, 3, "3"},
			{4, 1, "1"},
			{5, 2, "2"},
		},
		5,
		MaxHeap,
	}
	actualTest01 := BuildMaxHeap([]int{1, 2, 3, 4, 5}, []interface{}{"1", "2", "3", "4", "5"})
	require.Equal(t, expectedTest01, actualTest01)

	/* Test 02 right order */
	expectedTest02 := [][]element{
		{ /* 1 */
			{0, 0, nil},
			{1, 5, "5"},
			{2, 4, "4"},
			{3, 3, "3"},
			{4, 2, "2"},
			{5, 1, "1"},
		},
		{ /* 2 */
			{0, 0, nil},
			{1, 5, "5"},
			{2, 4, "4"},
			{3, 3, "3"},
			{4, 1, "1"},
			{5, 2, "2"},
		},
		{ /* 3 */
			{0, 0, nil},
			{1, 1, "1"},
			{2, 1, "2"},
			{3, 1, "3"},
			{4, 1, "4"},
			{5, 1, "5"},
		},
		{ /* 4 */
			{0, 0, nil},
			{1, 5, "5"},
			{2, 3, "3"},
			{3, 4, "4"},
			{4, 2, "2"},
			{5, 1, "1"},
		},
		{ /* 5 */
			{0, 0, nil},
			{1, 6, "6"},
			{2, 3, "3"},
			{3, 5, "5"},
			{4, 2, "2"},
			{5, 1, "1"},
			{6, 4, "4"},
		},
		{ /* 6 */
			{0, 0, nil},
			{1, 100, nil},
			{2, 67, nil},
			{3, 13, nil},
			{4, 15, nil},
			{5, 3, nil},
			{6, 5, nil},
			{7, 4, nil},
			{8, 6, nil},
			{9, 4, nil},
			{10, 2, nil},
			{11, 0, nil},
			{12, -11, nil},
			{13, 5, nil},
			{14, 4, nil},
			{15, 1, nil},
		},
	}

	testCasesTest02 := []myTestBuildHeap{
		{ /* 1 */ // no changes
			[]int{5, 4, 3, 2, 1},
			[]interface{}{"5", "4", "3", "2", "1"},
		},
		{ /* 2 */ // reversed
			[]int{1, 2, 3, 4, 5},
			[]interface{}{"1", "2", "3", "4", "5"},
		},
		{ /* 3 */ // equals
			[]int{1, 1, 1, 1, 1},
			[]interface{}{"1", "2", "3", "4", "5"},
		},
		{ /* 4 */ // one swap
			[]int{5, 2, 4, 3, 1},
			[]interface{}{"5", "2", "4", "3", "1"},
		},
		{ /* 5 */ // sifting from down to up
			[]int{5, 3, 4, 2, 1, 6},
			[]interface{}{"5", "3", "4", "2", "1", "6"},
		},
		{ /* 6 */ // random
			[]int{5, 3, 100, 4, 2, 5, 1, 6, 15, 67, 0, -11, 13, 4, 4},
			[]interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
	}

	for i := 0; i < len(testCasesTest02); i++ {
		testCase := testCasesTest02[i]
		actual := BuildMaxHeap(testCase.priority, testCase.values)
		require.Equal(t, expectedTest02[i], actual.arr, fmt.Sprintf("testCase: %v", i+1))
	}
}

func TestBuildMinHeap(t *testing.T) {
	/* Test 01 right format */
	expectedTest01 := &Heap{
		[]element{
			{0, 0, nil},
			{1, 2, "2"},
			{2, 3, "3"},
			{3, 5, "5"},
			{4, 4, "4"},
			{5, 6, "6"},
			{6, 7, "7"},
		},
		6,
		MinHeap,
	}
	actualTest01 := BuildMinHeap([]int{7, 6, 5, 4, 3, 2}, []interface{}{"7", "6", "5", "4", "3", "2"})
	require.Equal(t, expectedTest01, actualTest01)

	/* Test 02 right order */
	expectedTest02 := [][]element{
		{ /* 1 */
			{0, 0, nil},
			{1, 1, "1"},
			{2, 2, "2"},
			{3, 3, "3"},
			{4, 4, "4"},
			{5, 5, "5"},
		},
		{ /* 2 */
			{0, 0, nil},
			{1, 1, "1"},
			{2, 2, "2"},
			{3, 3, "3"},
			{4, 5, "5"},
			{5, 4, "4"},
		},
		{ /* 3 */
			{0, 0, nil},
			{1, 1, "1"},
			{2, 1, "2"},
			{3, 1, "3"},
			{4, 1, "4"},
			{5, 1, "5"},
		},
		{ /* 4 */
			{0, 0, nil},
			{1, 1, "1"},
			{2, 2, "2"},
			{3, 3, "3"},
			{4, 4, "4"},
			{5, 5, "5"},
		},
		{ /* 5 */
			{0, 0, nil},
			{1, 1, "1"},
			{2, 2, "2"},
			{3, 3, "3"},
			{4, 4, "4"},
			{5, 5, "5"},
			{6, 6, "6"},
		},
		{ /* 6 */
			{0, 0, nil},
			{1, -11, nil},
			{2, 0, nil},
			{3, 1, nil},
			{4, 4, nil},
			{5, 2, nil},
			{6, 5, nil},
			{7, 4, nil},
			{8, 6, nil},
			{9, 15, nil},
			{10, 67, nil},
			{11, 3, nil},
			{12, 100, nil},
			{13, 13, nil},
			{14, 4, nil},
			{15, 5, nil},
		},
	}

	testCasesTest02 := []myTestBuildHeap{
		{ /* 1 */ // no changes
			[]int{1, 2, 3, 4, 5},
			[]interface{}{"1", "2", "3", "4", "5"},
		},
		{ /* 2 */ // reversed
			[]int{5, 4, 3, 2, 1},
			[]interface{}{"5", "4", "3", "2", "1"},
		},
		{ /* 3 */ // equals
			[]int{1, 1, 1, 1, 1},
			[]interface{}{"1", "2", "3", "4", "5"},
		},
		{ /* 4 */ // one swap
			[]int{1, 4, 3, 2, 5},
			[]interface{}{"1", "4", "3", "2", "5"},
		},
		{ /* 5 */ // sifting from down to top
			[]int{5, 2, 3, 4, 1, 6},
			[]interface{}{"5", "2", "3", "4", "1", "6"},
		},
		{ /* 6 */ // random
			[]int{5, 3, 100, 4, 2, 5, 1, 6, 15, 67, 0, -11, 13, 4, 4},
			[]interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
	}

	for i := 0; i < len(testCasesTest02); i++ {
		testCase := testCasesTest02[i]
		actual := BuildMinHeap(testCase.priority, testCase.values)
		require.Equal(t, expectedTest02[i], actual.arr, fmt.Sprintf("testCase: %v", i+1))
	}
}


// ATTENTION!!!!
// Actual and expected in require are mixed up
// The author is squinting
func TestInsert(t *testing.T) {
	heap := New(MinHeap)

	/* 01 */
	heap.Insert(1, "string")
	require.Equal(t,
		[]int{heap.arr[0].priority, heap.arr[1].priority},
		[]int{0, 1},
		fmt.Sprintf("MinHeap 01\n"))
	require.Equal(t, heap.size, 1, fmt.Sprintf("MinHeap 01\n"))
	require.Equal(t,
		[]interface{}{heap.arr[0].value, heap.arr[1].value},
		[]interface{}{nil, "string"},
		fmt.Sprintf("MinHeap 01\n"))

	/* 02 */
	heap.Insert(42, 42)
	require.Equal(t,
		[]int{heap.arr[0].priority, heap.arr[1].priority, heap.arr[2].priority},
		[]int{0, 1, 42},
		fmt.Sprintf("MinHeap 02\n"))
	require.Equal(t, heap.size, 2, fmt.Sprintf("MinHeap 02\n"))
	require.Equal(t,
		[]interface{}{heap.arr[0].value, heap.arr[1].value, heap.arr[2].value},
		[]interface{}{nil, "string", 42},
		fmt.Sprintf("MinHeap 02\n"))

	/* 03 */
	heap.Insert(0, nil)
	require.Equal(t,
		[]int{heap.arr[0].priority, heap.arr[1].priority, heap.arr[2].priority, heap.arr[3].priority},
		[]int{0, 0, 42, 1},
		fmt.Sprintf("MinHeap 03\n"))
	require.Equal(t, heap.size, 3, fmt.Sprintf("MinHeap 03\n"))
	require.Equal(t,
		[]interface{}{heap.arr[0].value, heap.arr[1].value, heap.arr[2].value, heap.arr[3].value},
		[]interface{}{nil, nil, 42, "string"}, fmt.Sprintf("MinHeap 03\n"))


	heap = New(MaxHeap)

	/* 01 */
	heap.Insert(1, "string")
	require.Equal(t,
		[]int{heap.arr[0].priority, heap.arr[1].priority},
		[]int{0, 1},
		fmt.Sprintf("MaxHeap 01\n"))
	require.Equal(t, heap.size, 1, fmt.Sprintf("MaxHeap 01\n"))
	require.Equal(t,
		[]interface{}{heap.arr[0].value, heap.arr[1].value},
		[]interface{}{nil, "string"},
		fmt.Sprintf("MaxHeap 01\n"))

	/* 02 */
	heap.Insert(42, 42)
	require.Equal(t,
		[]int{heap.arr[0].priority, heap.arr[1].priority, heap.arr[2].priority},
		[]int{0, 42, 1},
		fmt.Sprintf("MaxHeap 02\n"))
	require.Equal(t, heap.size, 2, fmt.Sprintf("MaxHeap 02\n"))
	require.Equal(t,
		[]interface{}{heap.arr[0].value, heap.arr[1].value, heap.arr[2].value},
		[]interface{}{nil, 42, "string"},
		fmt.Sprintf("MaxHeap 02\n"))

	/* 03 */
	heap.Insert(0, nil)
	require.Equal(t,
		[]int{heap.arr[0].priority, heap.arr[1].priority, heap.arr[2].priority, heap.arr[3].priority},
		[]int{0, 42, 1, 0},
		fmt.Sprintf("MaxHeap 03\n"))
	require.Equal(t, heap.size, 3, fmt.Sprintf("MaxHeap 03\n"))
	require.Equal(t,
		[]interface{}{heap.arr[0].value, heap.arr[1].value, heap.arr[2].value, heap.arr[3].value},
		[]interface{}{nil, 42, "string", nil}, fmt.Sprintf("MaxHeap 03\n"))
}

// ATTENTION!!!!
// Actual and expected in require are mixed up
// The author is squinting
func TestExtractMin(t *testing.T) {

	/* 01 */
	heap := BuildMinHeap(
		[]int{13, 74, 38, 91, 0, 1, 1, 1, 4, 7, 3, 22, 566, 58},
		[]interface{}{"13", "74", "38", "91", "0", "1", "1", "1", "4", "7", "3", "22", "566", "58"})

	actualV, actualP := heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"0", 0}, fmt.Sprintf("Test01.1\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test01.2\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test01.3\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test01.4\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"3", 3}, fmt.Sprintf("Test01.5\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"4", 4}, fmt.Sprintf("Test01.6\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"7", 7}, fmt.Sprintf("Test01.7\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"13", 13}, fmt.Sprintf("Test01.8\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"22", 22}, fmt.Sprintf("Test01.9\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"38", 38}, fmt.Sprintf("Test01.10\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"58", 58}, fmt.Sprintf("Test01.11\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"74", 74}, fmt.Sprintf("Test01.12\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"91", 91}, fmt.Sprintf("Test01.13\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"566", 566}, fmt.Sprintf("Test01.14\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{nil, 0}, fmt.Sprintf("Test01.15\n"))

	/* 02 */ // MaxHeap
	heap = BuildMaxHeap(
		[]int{13, 74, 38, 91, 0, 1, 1, 1, 4, 7, 3, 22, 566, 58},
		[]interface{}{"13", "74", "38", "91", "0", "1", "1", "1", "4", "7", "3", "22", "566", "58"})

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"0", 0}, fmt.Sprintf("Test02.1\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test02.2\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test02.3\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test02.4\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"3", 3}, fmt.Sprintf("Test02.5\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"4", 4}, fmt.Sprintf("Test02.6\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"7", 7}, fmt.Sprintf("Test02.7\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"13", 13}, fmt.Sprintf("Test02.8\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"22", 22}, fmt.Sprintf("Test02.9\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"38", 38}, fmt.Sprintf("Test02.10\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"58", 58}, fmt.Sprintf("Test02.11\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"74", 74}, fmt.Sprintf("Test02.12\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"91", 91}, fmt.Sprintf("Test02.13\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"566", 566}, fmt.Sprintf("Test02.14\n"))

	actualV, actualP = heap.ExtractMin()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{nil, 0}, fmt.Sprintf("Test02.15\n"))
}

// ATTENTION!!!!
// Actual and expected in require are mixed up
// The author is squinting
func TestExtractMax(t *testing.T) {

	/* 01 */
	heap := BuildMaxHeap(
		[]int{13, 74, 38, 91, 0, 1, 1, 1, 4, 7, 3, 22, 566, 58},
		[]interface{}{"13", "74", "38", "91", "0", "1", "1", "1", "4", "7", "3", "22", "566", "58"})

	actualV, actualP := heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"566", 566}, fmt.Sprintf("Test01.1\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"91", 91}, fmt.Sprintf("Test01.2\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"74", 74}, fmt.Sprintf("Test01.3\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"58", 58}, fmt.Sprintf("Test01.4\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"38", 38}, fmt.Sprintf("Test01.5\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"22", 22}, fmt.Sprintf("Test01.6\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"13", 13}, fmt.Sprintf("Test01.7\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"7", 7}, fmt.Sprintf("Test01.8\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"4", 4}, fmt.Sprintf("Test01.9\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"3", 3}, fmt.Sprintf("Test01.10\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test01.11\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test01.12\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test01.13\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"0", 0}, fmt.Sprintf("Test01.14\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{nil, 0}, fmt.Sprintf("Test01.15\n"))

	/* 02 */ // MinHeap
	heap = BuildMinHeap(
		[]int{13, 74, 38, 91, 0, 1, 1, 1, 4, 7, 3, 22, 566, 58},
		[]interface{}{"13", "74", "38", "91", "0", "1", "1", "1", "4", "7", "3", "22", "566", "58"})

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"566", 566}, fmt.Sprintf("Test02.1\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"91", 91}, fmt.Sprintf("Test02.2\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"74", 74}, fmt.Sprintf("Test02.3\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"58", 58}, fmt.Sprintf("Test02.4\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"38", 38}, fmt.Sprintf("Test02.5\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"22", 22}, fmt.Sprintf("Test02.6\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"13", 13}, fmt.Sprintf("Test02.7\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"7", 7}, fmt.Sprintf("Test02.8\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"4", 4}, fmt.Sprintf("Test02.9\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"3", 3}, fmt.Sprintf("Test02.10\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test02.11\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test02.12\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"1", 1}, fmt.Sprintf("Test02.13\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{"0", 0}, fmt.Sprintf("Test02.14\n"))

	actualV, actualP = heap.ExtractMax()
	require.Equal(t, []interface{}{actualV, actualP}, []interface{}{nil, 0}, fmt.Sprintf("Test02.15\n"))
}