package iterator

type Aggregate interface {
	CreateIterator() Iterator
}
