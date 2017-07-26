package priorityqueue

import "testing"

// myKey implements Key
type myKey int

func (d1 myKey) Less(d Key) bool {
	d2 := d.(myKey)
	return d1-d2 < 0
}

type Person struct {
	Name string
	Age  int
}

func (d1 Person) Less(d Key) bool {
	d2 := d.(Person)
	return d1.Age-d2.Age < 0
}

var testCases = struct {
	in []myKey
	ex []myKey
}{
	[]myKey{1, 7, 7, 19, 1, 18, 5, 0, 16, 0, 14, 11, 2, 9, 8, 14, 11, 5, 17, 6},
	[]myKey{0, 0, 1, 1, 2, 5, 5, 6, 7, 7, 8, 9, 11, 11, 14, 14, 16, 17, 18, 19},
}

func testClient(minPQ MinPQ, t *testing.T) {
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
	for _, v := range testCases.in {
		minPQ.Insert(v)
	}
	for _, v := range testCases.ex {
		min := minPQ.Remove()
		if min != v {
			t.Errorf("Expected %v, Got %v", v, min)
		}

	}
}
