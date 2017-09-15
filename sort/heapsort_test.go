package sort

import (
	"testing"

	pq "github.com/hassansin/algorithms/priorityqueue"
)

// myKey implements Key
type myKey int

func (d1 myKey) Less(d pq.Key) bool {
	d2 := d.(myKey)
	return d1-d2 < 0
}

var testCases = struct {
	in []int
	ex []int
}{
	[]int{1, 7, 7, 19, 1, 18, 5, 0, 16, 0, 14, 11, 2, 9, 8, 14, 11, 5, 17, 6},
	[]int{0, 0, 1, 1, 2, 5, 5, 6, 7, 7, 8, 9, 11, 11, 14, 14, 16, 17, 18, 19},
}

func TestHeapSort(t *testing.T) {
	data := make([]pq.Key, len(testCases.in))

	for i := range testCases.in {
		data[i] = myKey(testCases.in[i])
	}
	Heapsort(data)
	for i := range testCases.ex {
		if data[i] != myKey(testCases.ex[i]) {
			t.Errorf("Expected %v, got %v", testCases.ex[i], data[i])
		}
	}
}
