package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// TODO 使っていない
type Aggregate interface {
	Iterator() Iterator
}
