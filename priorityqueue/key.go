package priorityqueue

type Key interface {
	// Less checks whether the current Key is smaller than the given Key
	Less(Key) bool
}
