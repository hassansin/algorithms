package sort

import pq "github.com/hassansin/algorithms/priorityqueue"

func sink(pq []pq.Key, i, l int) {
	for 2*i+1 < l {
		j := 2*i + 1
		if j+1 < l && pq[j].Less(pq[j+1]) {
			j++
		}
		if pq[j].Less(pq[i]) {
			return
		}
		swap(pq, i, j)
		i = j
	}
}
func swap(pq []pq.Key, i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func heapify(pq []pq.Key) {
	l := len(pq)
	for i := (l - 1) / 2; i >= 0; i-- {
		sink(pq, i, l)
	}
}

// Heapsort sorts []Key using Heap ordered priority queue
// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func Heapsort(pq []pq.Key) {
	heapify(pq)
	l := len(pq)
	for i := 0; i < l; i++ {
		swap(pq, 0, l-i-1)
		sink(pq, 0, l-i-1)
	}
}
