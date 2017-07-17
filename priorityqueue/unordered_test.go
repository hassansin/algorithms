package priorityqueue

import "testing"

// myKey implements Key
type myKey int

func (d1 myKey) Less(d Key) bool {
	d2 := d.(myKey)
	return d1-d2 < 0
}

func TestUnorderedMinPQ(t *testing.T) {
	minPQ := NewUnorderedMinPQ()
	for i := 0; i < 20; i++ {
		minPQ.Insert(myKey(i))
	}
	if minPQ.Size() != 20 {
		t.Errorf("Expected size of %v, Got %v", 20, minPQ.Size())
	}
	if minPQ.IsEmpty() {
		t.Errorf("Expected non-empty, got empty")
	}
	for i := 0; i < 20; i++ {
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
