package priorityqueue

type UnorderedMinPQ []Key

func (pq *UnorderedMinPQ) Insert(k Key) {
	*pq = append(*pq, k)
}

func (pq *UnorderedMinPQ) Remove() Key {
	key, idx := pq.Peek()
	pq.swap(idx, len(*pq)-1)
	*pq = (*pq)[:len(*pq)-1]
	return key
}

func (pq *UnorderedMinPQ) Peek() (Key, int) {
	idx := 0
	for i, k := range *pq {
		if k.Less((*pq)[idx]) {
			idx = i
		}
	}
	return (*pq)[idx], idx
}

func (pq UnorderedMinPQ) Size() int {
	return len(pq)
}

func (pq UnorderedMinPQ) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq UnorderedMinPQ) swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
