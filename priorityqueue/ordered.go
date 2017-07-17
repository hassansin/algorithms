package priorityqueue

type orderedMinPQ []Key

func NewOrderedMinPQ() MinPQ {
	return &orderedMinPQ{}
}

// Insert inserts a new Key into the queue. Time Complexity: O(n)
func (pq *orderedMinPQ) Insert(k Key) {
	*pq = append(*pq, k)
	idx := len(*pq) - 1
	for i, k := range *pq {
		if k.Less((*pq)[idx]) {
			pq.swap(idx, i)
		}
	}
}

// Remove removes the minimum Key from the queue and returns it. Time Complexity: O(1)
func (pq *orderedMinPQ) Remove() Key {
	key, pqd := (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
	*pq = pqd
	return key
}

// Peek returns the minimum Key in the queue. Time Complexity: O(1)
func (pq orderedMinPQ) Peek() Key {
	return pq[len(pq)-1]
}

// Size returns the size of the queue
func (pq orderedMinPQ) Size() int {
	return len(pq)
}

// IsEmpty checks whether the queue is empty or not
func (pq orderedMinPQ) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq orderedMinPQ) swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
