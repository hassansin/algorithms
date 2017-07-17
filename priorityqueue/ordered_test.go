package priorityqueue

import "testing"

func TestOrderedMinPQ(t *testing.T) {
	minPQ := NewOrderedMinPQ()
	testClient(minPQ, t)
}
