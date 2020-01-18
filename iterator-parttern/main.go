package main

import (
	"fmt"

	"github.com/liquid-dev/design-patterns-in-go/iterator-partern/book"
)

func main() {
	bs := book.NewShelf(4)
	bs.AppendBook(book.New("Around the World in 80 Days"))
	bs.AppendBook(book.New("Bible"))
	bs.AppendBook(book.New("Cinderella"))
	bs.AppendBook(book.New("Daddy-Long-Legs"))
	it := bs.Iterator()
	for it.HasNext() {
		b := it.Next().(*book.Book)
		fmt.Println(b.GetName())
	}
}
