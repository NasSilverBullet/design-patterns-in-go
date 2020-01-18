package book

import (
	"github.com/liquid-dev/design-patterns-in-go/iterator-partern/iterator"
)

type BookShelf []*Book

func NewShelf(maxsize int) BookShelf {
	return make(BookShelf, 0, maxsize)
}

func (bs BookShelf) GetBookAt(index int) *Book {
	return bs[index]
}

func (bs *BookShelf) AppendBook(b *Book) {
	*bs = append(*bs, b)

}

func (bs BookShelf) GetLength() int {
	return len(bs)
}

func (bs BookShelf) Iterator() iterator.Iterator {
	return NewShelfIterator(bs)
}
