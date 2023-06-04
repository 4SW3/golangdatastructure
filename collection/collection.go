package collection

// base interface that all data structures implement
type Collection interface {
	Len() int
	Repr() string
	// Clear()
}
