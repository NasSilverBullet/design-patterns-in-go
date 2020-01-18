package book_test

import (
	"reflect"
	"testing"

	"github.com/liquid-dev/design-patterns-in-go/iterator-pattern/book"
)

func TestBookShelf_GetBookAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		bs   book.BookShelf
		args args
		want *book.Book
	}{
		{"Head", append(book.NewShelf(1), book.New("hoge")), args{index: 0}, book.New("hoge")},
		{"Second", append(append(book.NewShelf(2), book.New("hoge")), book.New("huga")), args{index: 1}, book.New("huga")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.GetBookAt(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookShelf.GetBookAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
