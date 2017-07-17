package priorityqueue

import "testing"

func TestBinaryHeapMinPQ(t *testing.T) {
	minPQ := NewBinaryHeapMinPQ()
	for i := 20; i > 0; i-- {
		minPQ.Insert(myKey(i))
	}
	if minPQ.Size() != 20 {
		t.Errorf("Expected size of %v, Got %v", 20, minPQ.Size())
	}
	if minPQ.IsEmpty() {
		t.Errorf("Expected non-empty, got empty")
	}

	for i := 1; i <= 20; i++ {
		min := minPQ.Peek()
		if min != myKey(i) {
			t.Errorf("Expected %v, Got %v", i, min)
		}
		min2 := minPQ.Remove()
		if min2 != min {
			t.Errorf("Expected %v, Got %v", min, min2)
		}
	}
	if !minPQ.IsEmpty() {
		t.Errorf("Expected empty, got non-empty")
	}
}
