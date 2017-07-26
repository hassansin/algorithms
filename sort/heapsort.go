package sort

import pq "github.com/hassansin/algorithms/priorityqueue"

// sink implements the heap priority on pq[i,j].
func sink(data []pq.Key, i, l int) {
	for 2*i+1 < l {
		j := 2*i + 1
		if j+1 < l && data[j].Less(data[j+1]) {
			j++
		}
		if data[j].Less(data[i]) {
			return
		}
		swap(data, i, j)
		i = j
	}
}

// swap swaps elements with indexes i and j.
func swap(data []pq.Key, i, j int) {
	data[i], data[j] = data[j], data[i]
}
func heapify(data []pq.Key) {
	l := len(data)
	for i := (l - 1) / 2; i >= 0; i-- {
		sink(data, i, l)
	}
}

// Heapsort sorts []Key using Heap ordered priority queue.
// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func Heapsort(data []pq.Key) {
	heapify(data)
	l := len(data)
	for i := 0; i < l; i++ {
		swap(data, 0, l-i-1)
		sink(data, 0, l-i-1)
	}
}
