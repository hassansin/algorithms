package priorityqueue

type BinaryHeapMinPQ []Key

func (pq *BinaryHeapMinPQ) Insert(k Key) {
	if len(*pq) == 0 {
		*pq = append(*pq, nil)
	}
	*pq = append(*pq, k)
	pq.swim(len(*pq) - 1)
}

func (pq *BinaryHeapMinPQ) Remove() Key {
	pq.swap(1, len(*pq)-1)
	key, pqp := (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
	*pq = pqp
	pq.sink(1)
	return key
}

func (pq BinaryHeapMinPQ) Peek() Key {
	return pq[1]
}

func (pq BinaryHeapMinPQ) Size() int {
	if len(pq) == 0 {
		return 0
	}
	return len(pq) - 1
}

func (pq BinaryHeapMinPQ) IsEmpty() bool {
	return pq.Size() <= 1
}

func (pq BinaryHeapMinPQ) swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq BinaryHeapMinPQ) swim(i int) {
	for i > 1 && pq[i].Less(pq[i/2]) {
		pq.swap(i, i/2)
		i = i / 2
	}
}

func (pq BinaryHeapMinPQ) sink(i int) {
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
