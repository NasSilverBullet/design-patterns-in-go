package chardisplay

import "fmt"

type CharDisplay struct {
	ch string
}

func New(c string) *CharDisplay {
	return &CharDisplay{
		ch: c,
	}
}

func (cd *CharDisplay) Open() {
	fmt.Print("<<")
}

func (cd *CharDisplay) Print() {
	fmt.Print(cd.ch)
}

func (cd *CharDisplay) Close() {
	fmt.Println(">>")
}
