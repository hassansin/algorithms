package priorityqueue

import "testing"

func TestBinaryHeapMinPQ(t *testing.T) {
	minPQ := NewBinaryHeapMinPQ()
	testClient(minPQ, t)
}
