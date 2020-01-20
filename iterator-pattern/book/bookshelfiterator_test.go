package book_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/iterator-pattern/book"
)

func TestBookShelfIterator_HasNext(t *testing.T) {
	type fields struct {
		bookShelf book.BookShelf
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"True", fields{bookShelf: append(book.NewShelf(1), book.New("hoge"))}, true},
		{"False", fields{bookShelf: book.NewShelf(0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := book.NewShelfIterator(tt.fields.bookShelf)
			if got := bs.HasNext(); got != tt.want {
				t.Errorf("book.BookShelfIterator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookShelfIterator_Next(t *testing.T) {
	type fields struct {
		bookShelf book.BookShelf
		index     int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"Success", fields{bookShelf: append(book.NewShelf(1), book.New("hoge"))}, book.New("hoge")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := book.NewShelfIterator(tt.fields.bookShelf)
			if got := bs.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("book.BookShelfIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
