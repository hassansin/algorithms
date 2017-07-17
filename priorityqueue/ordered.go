package priorityqueue

type OrderedMinPQ []Key

func (pq *OrderedMinPQ) Insert(k Key) {
	*pq = append(*pq, k)
	idx := len(*pq) - 1
	for i, k := range *pq {
		if k.Less((*pq)[idx]) {
			pq.swap(idx, i)
		}
	}
}

func (pq *OrderedMinPQ) Remove() Key {
	key, pqd := (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
	*pq = pqd
	return key
}

func (pq OrderedMinPQ) Peek() Key {
	return pq[len(pq)-1]
}

func (pq OrderedMinPQ) Size() int {
	return len(pq)
}

func (pq OrderedMinPQ) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq OrderedMinPQ) swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
