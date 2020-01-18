package book

import "github.com/liquid-dev/design-patterns-in-go/iterator-pattern/iterator"

type BookShelf []*Book

func NewShelf(maxsize int) BookShelf {
	return make(BookShelf, 0, maxsize)
}

func (bs BookShelf) GetBookAt(index int) *Book {
	return bs[index]
}

func (bs BookShelf) Iterator() iterator.Iterator {
	return NewShelfIterator(bs)
}
