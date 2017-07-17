package priorityqueue

type unorderedMinPQ []Key

func NewUnorderedMinPQ() MinPQ {
	return &unorderedMinPQ{}
}

// Insert inserts a new Key into the queue. Time Complexity: O(1)
func (pq *unorderedMinPQ) Insert(k Key) {
	*pq = append(*pq, k)
}

// Remove removes the minimum Key from the queue and returns it. Time Complexity: O(n)
func (pq *unorderedMinPQ) Remove() Key {
	idx, key := pq.peek()
	pq.swap(idx, len(*pq)-1)
	*pq = (*pq)[:len(*pq)-1]
	return key
}

// Peek returns the minimum Key in the queue. Time Complexity: O(n)
func (pq unorderedMinPQ) Peek() Key {
	_, key := pq.peek()
	return key
}

func (pq *unorderedMinPQ) peek() (int, Key) {
	idx := 0
	for i, k := range *pq {
		if k.Less((*pq)[idx]) {
			idx = i
		}
	}
	return idx, (*pq)[idx]
}

// Size returns the size of the queue
func (pq unorderedMinPQ) Size() int {
	return len(pq)
}

// IsEmpty checks whether the queue is empty or not
func (pq unorderedMinPQ) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq unorderedMinPQ) swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
