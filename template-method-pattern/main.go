package main

import (
	"github.com/NasSilverBullet/design-patterns-in-go/template-method-pattern/abstractdisplay"
	"github.com/NasSilverBullet/design-patterns-in-go/template-method-pattern/chardisplay"
	"github.com/NasSilverBullet/design-patterns-in-go/template-method-pattern/stringdisplay"
)

func main() {
	d1 := abstractdisplay.New(chardisplay.New("H"))
	d2 := abstractdisplay.New(stringdisplay.New("Hello, world."))
	d3 := abstractdisplay.New(stringdisplay.New("Good Bye!!"))

	d1.Display()
	d2.Display()
	d3.Display()
}
