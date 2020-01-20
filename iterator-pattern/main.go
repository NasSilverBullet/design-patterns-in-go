package main

import (
	"fmt"

	"github.com/NasSilverBullet/design-patterns-in-go/iterator-pattern/book"
)

func main() {
	bs := book.NewShelf(4)
	bs = append(bs, book.New("Around the World in 80 Days"))
	bs = append(bs, book.New("Bible"))
	bs = append(bs, book.New("Cinderella"))
	bs = append(bs, book.New("Daddy-Long-Legs"))
	it := bs.Iterator()
	for it.HasNext() {
		b := it.Next().(*book.Book)
		fmt.Println(b.GetName())
	}
}
