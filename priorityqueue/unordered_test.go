package priorityqueue

import "testing"

func TestUnorderedMinPQ(t *testing.T) {
	minPQ := NewUnorderedMinPQ()
	testClient(minPQ, t)
}
