package stringdisplay

import (
	"fmt"
)

type StringDisplay struct {
	string string
	width  int
}

func New(s string) *StringDisplay {
	return &StringDisplay{
		string: s,
		width:  len(s),
	}
}

func (sd *StringDisplay) Open() {
	sd.PrintLine()
}

func (sd *StringDisplay) Print() {
	fmt.Printf("|%s|\n", sd.string)
}

func (sd *StringDisplay) Close() {
	sd.PrintLine()
}

func (sd *StringDisplay) PrintLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
