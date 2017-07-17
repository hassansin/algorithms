package priorityqueue

type binaryHeapMinPQ []Key

func NewBinaryHeapMinPQ() MinPQ {
	return &binaryHeapMinPQ{}
}

// Insert inserts the key into the heap. Time Complexity: O(log(n))
func (pq *binaryHeapMinPQ) Insert(k Key) {
	if len(*pq) == 0 {
		*pq = append(*pq, nil)
	}
	*pq = append(*pq, k)
	pq.swim(len(*pq) - 1)
}

// Remove removes the minimum key from the heap and returns it. Time Complexity: O(log(n))
func (pq *binaryHeapMinPQ) Remove() Key {
	pq.swap(1, len(*pq)-1)
	key, pqp := (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
	*pq = pqp
	pq.sink(1)
	return key
}

// Peek returns the minimum key from the heap. Time Complexity: O(1)
func (pq binaryHeapMinPQ) Peek() Key {
	return pq[1]
}

// Size returns the heap size.
func (pq binaryHeapMinPQ) Size() int {
	if len(pq) == 0 {
		return 0
	}
	return len(pq) - 1
}

// IsEmpty checks if the heap is empty or not
func (pq binaryHeapMinPQ) IsEmpty() bool {
	return pq.Size() <= 1
}

func (pq binaryHeapMinPQ) swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq binaryHeapMinPQ) swim(i int) {
	for i > 1 && pq[i].Less(pq[i/2]) {
		pq.swap(i, i/2)
		i = i / 2
	}
}

func (pq binaryHeapMinPQ) sink(i int) {
	for 2*i < len(pq) {
		j := 2 * i
		if j+1 < len(pq) && pq[j+1].Less(pq[j]) {
			j++
		}
		if pq[i].Less(pq[j]) {
			break
		}
		pq.swap(i, j)
		i = j
	}
}
