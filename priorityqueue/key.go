package priorityqueue

type Key interface {
	Less(Key) bool
}
