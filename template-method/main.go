package main

import (
	"github.com/liquid-dev/design-patterns-in-go/template-method/abstractdisplay"
	"github.com/liquid-dev/design-patterns-in-go/template-method/chardisplay"
)

func main() {
	d1 := abstractdisplay.New(chardisplay.New("H"))
	d2 := abstractdisplay.New(chardisplay.New("Hello, world."))
	d3 := abstractdisplay.New(chardisplay.New("Good Bye!!"))

	d1.Display()
	d2.Display()
	d3.Display()
}
