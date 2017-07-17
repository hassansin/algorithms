package priorityqueue

type MinPQ interface {
	// Insert inserts a new Key into the collection
	Insert(Key)
	// Remove removes the minimum Key from the collection and returns it
	Remove() Key
	// Peek returns the minimum Key from the collection
	Peek() Key
	// Size returns the size of the collection
	Size() int
	// IsEmpty checks whether the collection is empty or not
	IsEmpty() bool
}
