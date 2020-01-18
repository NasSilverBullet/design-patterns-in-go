package book

type BookShelfIterator struct {
	bookShelf BookShelf
	index     int
}

func NewShelfIterator(bs BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bs,
	}
}

func (bs *BookShelfIterator) HasNext() bool {
	if bs.index < len(bs.bookShelf) {
		return true
	}
	return false
}

func (bs *BookShelfIterator) Next() interface{} {
	b := bs.bookShelf.GetBookAt(bs.index)
	bs.index++
	return b
}
